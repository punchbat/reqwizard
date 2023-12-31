package repository

import (
	"context"
	"errors"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/auth/repository/dto"
	pkgGorm "reqwizard/pkg/postgres/gorm"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *pkgGorm.Gorm
}

func NewRepository(db *pkgGorm.Gorm) *Repository {
	return &Repository{
		db: db,
	}
}

func (r Repository) CreateUser(ctx context.Context, req *domain.User) error {
	query := r.db.Conn

	newUser := dto.ConvertUserFromDomain(req)

	err := query.Create(&newUser).Error
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) UpdateUser(ctx context.Context, req *domain.User) error {
	query := r.db.Conn

	newUser := dto.ConvertUserFromDomain(req)

	err := query.Model(&dto.User{}).Where("id = ?", req.ID).Updates(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := r.db.Conn

	var userDTO dto.User

	err := query.Where("email = ?", email).First(&userDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return dto.ConvertUserToDomain(&userDTO), nil
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	query := r.db.Conn

	var userDTO dto.User

	userID := uuid.MustParse(id)

	err := query.Where("id = ?", userID).First(&userDTO).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	return dto.ConvertUserToDomain(&userDTO), nil
}

func (r *Repository) DeleteUser(ctx context.Context, id string) error {
	query := r.db.Conn

	userID := uuid.MustParse(id)

	err := query.Where("id = ?", userID).Delete(&dto.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUsers(ctx context.Context, ids []string) error {
	query := r.db.Conn

	var userIDs []uuid.UUID
	for _, id := range ids {
		userIDs = append(userIDs, uuid.MustParse(id))
	}

	err := query.Where("id IN ?", userIDs).Delete(&dto.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUnverifiedUsersCreatedBefore(ctx context.Context, before time.Time) ([]*domain.User, error) {
	query := r.db.Conn

	var userDTOs []*dto.User
	err := query.Where("verified = false AND created_at < ?", before).Find(&userDTOs).Error
	if err != nil {
		return nil, err
	}

	var unverifiedUsers []*domain.User
	for _, userDTO := range userDTOs {
		unverifiedUsers = append(unverifiedUsers, dto.ConvertUserToDomain(userDTO))
	}

	return unverifiedUsers, nil
}