package usecase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"reqwizard/internal/domain"
	"reqwizard/internal/services/email"
	service_email "reqwizard/internal/services/email"
	"reqwizard/internal/shared/utils"
	"strconv"
	"time"

	"reqwizard/internal/routes/auth"
	"reqwizard/internal/routes/role"
	"reqwizard/internal/routes/userRole"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *domain.User `json:"user"`
}

type UseCase struct {
	repo           auth.Repository
	roleRepo       role.Repository
	userRoleRepo   userRole.Repository
	mailer         *email.Mailer
	signingKey     []byte
	expireDuration time.Duration
}

func NewUseCase(
	repo auth.Repository,
	roleRepo role.Repository,
	userRoleRepo userRole.Repository,

	mailer *service_email.Mailer,
	signingKey []byte,
	tokenTTLHours time.Duration) *UseCase {
	return &UseCase{
		repo:         repo,
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,

		mailer:         mailer,
		signingKey:     signingKey,
		expireDuration: time.Hour * tokenTTLHours,
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswordHash(password1, password2 string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(password1), []byte(password2))
	if err != nil {
		return false, errors.New("invalid password")
	}
	return true, nil
}

func (uc *UseCase) MakeClearUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	user.Password = ""
	user.PasswordConfirm = ""

	userRoles, err := uc.userRoleRepo.GetUserRoles(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	roles, err := uc.roleRepo.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	var resUserRoles []*domain.UserRole
	for _, userRole := range userRoles {
		for _, role := range roles {
			if userRole.RoleID == role.ID {
				resUserRoles = append(resUserRoles, &domain.UserRole{
					Name:   domain.UserRoleName(role.Name),
					Status: userRole.Status,
				})

				break
			}
		}
	}

	user.UserRoles = resUserRoles

	return user, nil
}

func (uc *UseCase) SignUp(ctx context.Context, inp *auth.SignUpInput) error {
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err == nil {
		return auth.ErrUserIsExist
	}

	hashPassword, err := HashPassword(inp.Password)
	if err != nil {
		return err
	}

	hashPasswordConfirm, err := HashPassword(inp.PasswordConfirm)
	if err != nil {
		return err
	}

	// Находим роль обычного юзера
	role, err := uc.roleRepo.GetRoleByName(ctx, string(inp.Role))
	if err != nil {
		return err
	}

	// подумать
	userID := uuid.New()
	user = &domain.User{
		ID:              userID.String(),
		Email:           inp.Email,
		Password:        hashPassword,
		PasswordConfirm: hashPasswordConfirm,
		Verified:        false,
	}
	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	// Создаем юзер.роль
	userRoleID := uuid.New()
	userRole := domain.UserRole{
		ID:     userRoleID.String(),
		UserID: userID.String(),
		RoleID: role.ID,
		Status: domain.UserRoleStatusApproved,
	}
	err = uc.userRoleRepo.CreateUserRole(ctx, &userRole)
	if err != nil {
		return err
	}

	return nil
}

type EmailContent struct {
	VerifyCode string
}

func (uc *UseCase) SendVerifyCode(ctx context.Context, inp *auth.SendVerifyCodeInput) error {
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return auth.ErrUserNotFound
	}

	isEqual, err := ComparePasswordHash(user.Password, inp.Password)
	if err != nil {
		return err
	}
	if !isEqual {
		return auth.ErrEmailOrPassword
	}

	minVal := 100000
	maxVal := 999999
	randomCode := rand.Intn(maxVal-minVal+1) + minVal

	user.VerifyCode = strconv.Itoa(randomCode)
	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return err
	}

	// Отправляем письмо
	emailMessage := service_email.Message{
		Subject:      "Reqwizard: verify code",
		To:           []string{inp.Email},
		TemplateName: "VerifyCode",
		Content: EmailContent{
			VerifyCode: user.VerifyCode,
		},
	}
	uc.mailer.Send(&emailMessage)

	return nil
}

func (uc *UseCase) CheckVerifyCode(ctx context.Context, inp *auth.CheckVerifyCodeInput) (string, error) {
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return "", auth.ErrUserNotFound
	}

	isEqual, err := ComparePasswordHash(user.Password, inp.Password)
	if err != nil {
		return "", err
	}
	if !isEqual {
		return "", auth.ErrEmailOrPassword
	}

	if user.VerifyCode != inp.VerifyCode {
		return "", auth.ErrVerifyCodeNotMatch
	}

	user.Verified = true
	user.VerifyCode = ""

	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return "", err
	}

	return uc.GetToken(ctx, user)
}

func (uc *UseCase) SignIn(ctx context.Context, inp *auth.SignInInput) error {
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return auth.ErrUserNotFound
	}

	isEqual, err := ComparePasswordHash(user.Password, inp.Password)
	if err != nil {
		return err
	}
	if !isEqual {
		return auth.ErrEmailOrPassword
	}

	// Если юзер не verified, то он не может зайти
	if !user.Verified {
		return auth.ErrUserIsUnauthorized
	}

	return nil
}

func (uc *UseCase) GetToken(ctx context.Context, user *domain.User) (string, error) {
	user, err := uc.MakeClearUser(ctx, user)
	if err != nil {
		return "", err
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(uc.expireDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	completeSignedToken, err := token.SignedString(uc.signingKey)

	if err != nil {
		return "", err
	}

	return completeSignedToken, nil
}

func (uc *UseCase) ParseToken(ctx context.Context, accessToken string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return uc.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, auth.ErrInvalidAccessToken
}

func (uc *UseCase) GetProfile(ctx context.Context, inp *auth.GetProfileInput) (*domain.User, error) {
	user, err := uc.repo.GetUserByID(ctx, inp.ID)

	if err != nil {
		return nil, auth.ErrUserNotFound
	}

	user, err = uc.MakeClearUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UseCase) RemoveUnverifiedUsers(ctx context.Context) error {
	now := time.Now()
	interval := now.Add(-24 * time.Hour)

	// Получаем список неподтвержденных пользователей, созданных более 24 часов назад
	unverifiedUsers, err := uc.repo.GetUnverifiedUsersCreatedBefore(ctx, interval)
	if err != nil {
		return err
	}

	unverifiedUserIDs := utils.Map(unverifiedUsers, func(i *domain.User) string {
		return i.ID
	})

	if err = uc.repo.DeleteUsers(ctx, unverifiedUserIDs); err != nil {
		return err
	}

	return nil
}
