package usecase

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"reqwizard/internal/domain"
	"reqwizard/internal/services/email"
	service_email "reqwizard/internal/services/email"
	"reqwizard/internal/shared/utils"
	"strconv"
	"strings"
	"time"

	"reqwizard/internal/routes/auth"
	"reqwizard/internal/routes/role"
	"reqwizard/internal/routes/userRole"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *domain.User `json:"user"`
}

type UseCase struct {
	repo         auth.Repository
	roleRepo     role.Repository
	userRoleRepo userRole.Repository
	mailer       *email.Mailer
}

func NewUseCase(
	repo auth.Repository,
	roleRepo role.Repository,
	userRoleRepo userRole.Repository,

	mailer *service_email.Mailer) *UseCase {
	return &UseCase{
		repo:         repo,
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,

		mailer: mailer,
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

func (uc *UseCase) SignUp(ctx context.Context, inp *auth.SignUpInput) (int, error) {
	// * Юзер есть и не верифицирован
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err == nil && user.Verified == false {
		return 0, nil
	} else if err == nil && user.Verified == true {
		return http.StatusConflict, auth.ErrUserIsExist
	}

	hashPassword, err := HashPassword(inp.Password)
	if err != nil {
		return http.StatusNotAcceptable, err
	}

	hashPasswordConfirm, err := HashPassword(inp.PasswordConfirm)
	if err != nil {
		return http.StatusNotAcceptable, err
	}

	birthday, err := utils.GetTimeFromString(inp.Birthday)
	if err != nil {
		return http.StatusNotAcceptable, err
	}

	// * Создаем Аккаунт
	userID := uuid.New()
	user = &domain.User{
		ID:              userID.String(),
		Email:           inp.Email,
		Password:        hashPassword,
		PasswordConfirm: hashPasswordConfirm,
		Verified:        false,

		Name:     inp.Name,
		Surname:  inp.Surname,
		Gender:   domain.UserGender(inp.Gender),
		Birthday: birthday,
	}
	// * Проверяем наличие аватарки
	if inp.Avatar != nil {
		avatarExt := filepath.Ext(inp.AvatarName)
		avatarName := uuid.New().String() + avatarExt
		avatarPath := "uploads/avatars/" + avatarName

		// Чтение данных из inp.Avatar
		avatarData, err := ioutil.ReadAll(inp.Avatar)
		if err != nil {
			return http.StatusNotAcceptable, err
		}

		// Создание изображения из данных
		img, _, err := image.Decode(bytes.NewReader(avatarData))
		if err != nil {
			return http.StatusNotAcceptable, err
		}

		// Сохранить изображение в формате JPEG с заданным качеством (80%)
		outFile, err := os.Create(avatarPath)
		if err != nil {
			return http.StatusConflict, err
		}
		defer outFile.Close()

		format := strings.ToLower(avatarExt)
		switch format {
		case ".jpg", ".jpeg":
			// Сохранить изображение в формате JPEG с заданным качеством (80%)
			outFile, err := os.Create(avatarPath)
			if err != nil {
				return http.StatusConflict, err
			}
			defer outFile.Close()

			err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 80})
			if err != nil {
				return http.StatusConflict, err
			}

		case ".png":
			// Сохранить изображение в формате PNG
			outFile, err := os.Create(avatarPath)
			if err != nil {
				return http.StatusConflict, err
			}
			defer outFile.Close()

			err = png.Encode(outFile, img)
			if err != nil {
				return http.StatusConflict, err
			}

		default:
			return http.StatusUnsupportedMediaType, fmt.Errorf("неподдерживаемое расширение изображения: %s", format)
		}

		user.Avatar = avatarName
	}
	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return http.StatusConflict, err
	}

	defaultRoleEntity, err := uc.roleRepo.GetRoleByName(ctx, "user")
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defaultUserRoleEntity := domain.UserRole{
		ID:     uuid.New().String(),
		UserID: userID.String(),
		RoleID: defaultRoleEntity.ID,
		Status: domain.UserRoleStatusApproved,
	}
	err = uc.userRoleRepo.CreateUserRole(ctx, &defaultUserRoleEntity)
	if err != nil {
		return http.StatusConflict, err
	}

	if inp.Role != "user" {
		// Находим роль
		selectedRoleEntity, err := uc.roleRepo.GetRoleByName(ctx, string(inp.Role))
		if err != nil {
			return http.StatusInternalServerError, err
		}
		selectedUserRole := domain.UserRole{
			ID:     uuid.New().String(),
			UserID: userID.String(),
			RoleID: selectedRoleEntity.ID,
			Status: domain.UserRoleStatusApproved,
		}
		err = uc.userRoleRepo.CreateUserRole(ctx, &selectedUserRole)
		if err != nil {
			return http.StatusConflict, err
		}
	}

	return http.StatusCreated, nil
}

type EmailContent struct {
	VerifyCode string
}

func (uc *UseCase) SendVerifyCode(ctx context.Context, inp *auth.SendVerifyCodeInput) (int, error) {
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return http.StatusNotAcceptable, auth.ErrEmailOrPassword
	}

	isEqual, err := ComparePasswordHash(user.Password, inp.Password)
	if err != nil {
		return http.StatusNotAcceptable, auth.ErrUserNotFound
	}
	if !isEqual {
		return http.StatusNotAcceptable, auth.ErrEmailOrPassword
	}

	minVal := 100000
	maxVal := 999999
	randomCode := rand.Intn(maxVal-minVal+1) + minVal

	user.VerifyCode = strconv.Itoa(randomCode)
	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return http.StatusBadRequest, err
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

	return http.StatusOK, nil
}

func (uc *UseCase) CheckVerifyCode(ctx context.Context, inp *auth.CheckVerifyCodeInput) (string, int, error) {
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return "", http.StatusNotAcceptable, auth.ErrEmailOrPassword
	}

	isEqual, err := ComparePasswordHash(user.Password, inp.Password)
	if err != nil {
		return "", http.StatusNotAcceptable, auth.ErrEmailOrPassword
	}
	if !isEqual {
		return "", http.StatusNotAcceptable, auth.ErrEmailOrPassword
	}

	if user.VerifyCode != inp.VerifyCode {
		return "", http.StatusBadRequest, auth.ErrVerifyCodeNotMatch
	}

	user.Verified = true
	user.VerifyCode = ""

	if err := uc.repo.UpdateUser(ctx, user); err != nil {
		return "", http.StatusBadRequest, auth.ErrVerifyCodeNotMatch
	}

	return uc.GetToken(ctx, user)
}

func (uc *UseCase) SignIn(ctx context.Context, inp *auth.SignInInput) (int, error) {
	user, err := uc.repo.GetUserByEmail(ctx, inp.Email)
	if err != nil {
		return http.StatusNotAcceptable, auth.ErrEmailOrPassword
	}

	isEqual, err := ComparePasswordHash(user.Password, inp.Password)
	if err != nil {
		return http.StatusNotAcceptable, err
	}
	if !isEqual {
		return http.StatusNotAcceptable, auth.ErrEmailOrPassword
	}

	// Если юзер не verified, то он не может зайти
	if !user.Verified {
		return http.StatusUnauthorized, auth.ErrUserIsUnauthorized
	}

	return http.StatusOK, nil
}

func (uc *UseCase) GetToken(ctx context.Context, user *domain.User) (string, int, error) {
	user, err := uc.MakeClearUser(ctx, user)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * viper.GetDuration("auth.token.ttl"))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signingKey := []byte(viper.GetString("auth.signing_key"))

	completeSignedToken, err := token.SignedString(signingKey)

	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return completeSignedToken, http.StatusCreated, nil
}

func (uc *UseCase) GetProfile(ctx context.Context, id string) (*domain.User, int, error) {
	user, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, http.StatusNotFound, auth.ErrUserNotFound
	}

	user, err = uc.MakeClearUser(ctx, user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func (uc *UseCase) UpdateProfile(ctx context.Context, inp *auth.UpdateInput) (int, error) {
	user, err := uc.repo.GetUserByID(ctx, inp.ID)
	if err != nil {
		return http.StatusNotAcceptable, err
	}

	userRoles, err := uc.userRoleRepo.GetUserRoles(ctx, user.ID)

	newUserRoles := make(map[string]bool)
	for _, name := range inp.UserRoles {
		newUserRoles[name] = true
	}
	currentUserRoles := make(map[string]bool)
	for _, userRole := range userRoles {
		currentUserRoles[string(userRole.Name)] = true
	}

	// * Добавляем новые роли, которые есть в инпуте
	for _, userRoleName := range inp.UserRoles {
		if currentUserRoles[userRoleName] {
			continue
		}

		selectedRoleEntity, err := uc.roleRepo.GetRoleByName(ctx, string(userRoleName))
		if err != nil {
			return http.StatusInternalServerError, err
		}
		selectedUserRole := domain.UserRole{
			ID:     uuid.New().String(),
			UserID: user.ID,
			RoleID: selectedRoleEntity.ID,
			Status: domain.UserRoleStatusApproved,
		}
		err = uc.userRoleRepo.CreateUserRole(ctx, &selectedUserRole)
		if err != nil {
			return http.StatusConflict, err
		}
	}

	// * Удаляем роли, которых нет в инпуте
	for _, userRole := range userRoles {
		if newUserRoles[string(userRole.Name)] {
			continue
		}

		err = uc.userRoleRepo.DeleteUserRoleByID(ctx, userRole.ID)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	birthday, err := utils.GetTimeFromString(inp.Birthday)
	if err != nil {
		return http.StatusNotAcceptable, err
	}

	updateUser := &domain.User{
		ID:       user.ID,
		Name:     inp.Name,
		Surname:  inp.Surname,
		Gender:   domain.UserGender(inp.Gender),
		Birthday: birthday,
		Avatar:   user.Avatar,
	}

	// * Проверьте, есть ли у пользователя старая фотография, и удалите ее, если есть.
	if inp.AvatarName != "" && user.Avatar != "" {
		err := utils.DeleteFile("uploads/avatars/" + user.Avatar)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}
	if inp.Avatar != nil {
		avatarExt := filepath.Ext(inp.AvatarName)
		avatarName := uuid.New().String() + avatarExt
		avatarPath := "uploads/avatars/" + avatarName

		// Чтение данных из inp.Avatar
		avatarData, err := ioutil.ReadAll(inp.Avatar)
		if err != nil {
			return http.StatusNotAcceptable, err
		}

		// Создание изображения из данных
		img, _, err := image.Decode(bytes.NewReader(avatarData))
		if err != nil {
			return http.StatusNotAcceptable, err
		}

		// Сохранить изображение в формате JPEG с заданным качеством (80%)
		outFile, err := os.Create(avatarPath)
		if err != nil {
			return http.StatusConflict, err
		}
		defer outFile.Close()

		format := strings.ToLower(avatarExt)
		switch format {
		case ".jpg", ".jpeg":
			// Сохранить изображение в формате JPEG с заданным качеством (80%)
			outFile, err := os.Create(avatarPath)
			if err != nil {
				return http.StatusConflict, err
			}
			defer outFile.Close()

			err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 80})
			if err != nil {
				return http.StatusConflict, err
			}

		case ".png":
			// Сохранить изображение в формате PNG
			outFile, err := os.Create(avatarPath)
			if err != nil {
				return http.StatusConflict, err
			}
			defer outFile.Close()

			err = png.Encode(outFile, img)
			if err != nil {
				return http.StatusConflict, err
			}

		default:
			return http.StatusUnsupportedMediaType, fmt.Errorf("неподдерживаемое расширение изображения: %s", format)
		}

		updateUser.Avatar = avatarName
	}

	err = uc.repo.UpdateUser(ctx, updateUser)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (uc *UseCase) ParseToken(ctx context.Context, accessToken string) (*domain.User, int, error) {
	signingKey := []byte(viper.GetString("auth.signing_key"))

	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, http.StatusNotAcceptable, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, http.StatusOK, nil
	}

	return nil, http.StatusBadRequest, auth.ErrInvalidAccessToken
}

// * JOBS.
func (uc *UseCase) RemoveUnverifiedUsers(ctx context.Context) (int, error) {
	now := time.Now()
	interval := now.Add(-24 * time.Hour)

	// Получаем список неподтвержденных пользователей, созданных более 24 часов назад
	unverifiedUsers, err := uc.repo.GetUnverifiedUsersCreatedBefore(ctx, interval)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	unverifiedUserIDs := utils.Map(unverifiedUsers, func(i *domain.User) string {
		return i.ID
	})

	if err = uc.repo.DeleteUsers(ctx, unverifiedUserIDs); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
