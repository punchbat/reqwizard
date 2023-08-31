package repository

import (
	"context"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/userRole/repository/dto"
	"reqwizard/pkg/postgres/gorm"
)

type Repository struct {
	db *gorm.Gorm
}

func NewRepository(db *gorm.Gorm) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateUserRole(ctx context.Context, userRole *domain.UserRole) error {
	query := r.db.Conn

	model := dto.ConvertUserRoleFromDomain(userRole)
	err := query.Create(model).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserRoleByID(ctx context.Context, id string) (*domain.UserRole, error) {
	query := r.db.Conn

	var roleDTO dto.UserRole
	err := query.First(&roleDTO, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return dto.ConvertUserRoleToDomain(&roleDTO), nil
}

func (r *Repository) GetUserRoles(ctx context.Context, userId string) ([]*domain.UserRole, error) {
	query := r.db.Conn

	var rolesDTO []*dto.UserRole
	err := query.Find(&rolesDTO, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}

	var roles []*domain.UserRole
	for _, roleDTO := range rolesDTO {
		roles = append(roles, dto.ConvertUserRoleToDomain(roleDTO))
	}

	return roles, nil
}

func (r *Repository) DeleteUserRoleByID(ctx context.Context, id string) error {
	query := r.db.Conn

	err := query.Delete(&dto.UserRole{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateUserRoles(ctx context.Context, userRoles []*domain.UserRole) error {
	query := r.db.Conn

	for _, userRole := range userRoles {
		model := dto.ConvertUserRoleFromDomain(userRole)

		err := query.Model(&dto.UserRole{}).Where("id = ?", model.ID).Updates(model).Error
		if err != nil {
			return err
		}
	}

	return nil
}
