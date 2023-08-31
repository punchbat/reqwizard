package auth

import (
	"context"
	"reqwizard/internal/domain"
)

type Repository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
}

const CtxUserKey = "user"

type UseCase interface {
	SignUp(ctx context.Context, inp *SignUpInput) error
	SendVerifyCode(ctx context.Context, inp *SendVerifyCodeInput) error
	CheckVerifyCode(ctx context.Context, inp *CheckVerifyCodeInput) (string, error)

	SignIn(ctx context.Context, inp *SignInInput) error
	ParseToken(ctx context.Context, accessToken string) (*domain.User, error)

	GetProfile(ctx context.Context, inp *GetProfileInput) (*domain.User, error)
}