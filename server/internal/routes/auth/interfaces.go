package auth

import (
	"context"
	"reqwizard/internal/domain"
	"time"
)

type Repository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
	DeleteUsers(ctx context.Context, ids []string) error
	GetUnverifiedUsersCreatedBefore(ctx context.Context, before time.Time) ([]*domain.User, error)
}

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, inp *SignUpInput) (int, error)
	SendVerifyCode(ctx context.Context, inp *SendVerifyCodeInput) (int, error)
	CheckVerifyCode(ctx context.Context, inp *CheckVerifyCodeInput) (string, int, error)

	SignIn(ctx context.Context, inp *SignInInput) (int, error)
	ParseToken(ctx context.Context, accessToken string) (*domain.User, int, error)

	GetProfile(ctx context.Context, id string) (*domain.User, int, error)
	UpdateProfile(ctx context.Context, inp *UpdateInput) (int, error)

	RemoveUnverifiedUsers(ctx context.Context) (int, error)
}