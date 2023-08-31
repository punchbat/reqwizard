package repository

import (
	"context"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/role/repository/dto"
	"reqwizard/pkg/postgres/gorm"

	"github.com/google/uuid"
)

type Repository struct {
	db *gorm.Gorm
}

func NewRepository(db *gorm.Gorm) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetRoleByID(ctx context.Context, id string) (*domain.Role, error) {
	query := r.db.Conn

	uuid := uuid.MustParse(id)
	var roleDTO dto.Role

	err := query.Where("id = ?", uuid).First(&roleDTO).Error
	if err != nil {
		return nil, err
	}

	return dto.ConvertRoleToDomain(&roleDTO), nil
}

func (r *Repository) GetRoleByIDs(ctx context.Context, ids []string) ([]*domain.Role, error) {
	if len(ids) == 0 {
		return []*domain.Role{}, nil
	}

	query := r.db.Conn

	var roles []*domain.Role
	var uuids []uuid.UUID

	for _, id := range ids {
		uuids = append(uuids, uuid.MustParse(id))
	}

	err := query.Where("id IN (?)", uuids).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *Repository) GetRoleByName(ctx context.Context, name string) (*domain.Role, error) {
	query := r.db.Conn

	var roleDTO dto.Role

	err := query.Where("name = ?", name).First(&roleDTO).Error
	if err != nil {
		return nil, err
	}

	return dto.ConvertRoleToDomain(&roleDTO), nil
}

func (r *Repository) GetRoles(ctx context.Context) ([]*domain.Role, error) {
	query := r.db.Conn

	var rolesDTO []*dto.Role
	err := query.Find(&rolesDTO).Error
	if err != nil {
		return nil, err
	}

	var roles []*domain.Role
	for _, roleDTO := range rolesDTO {
		roles = append(roles, dto.ConvertRoleToDomain(roleDTO))
	}

	return roles, nil
}