package repository

import (
	"context"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/application"
	"reqwizard/internal/routes/application/repository/dto"
	"reqwizard/internal/shared/utils"
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

func (r Repository) CreateApplication(ctx context.Context, req *domain.Application) error {
	query := r.db.Conn

	newApplication := dto.ConvertApplicationFromDomain(req)

	err := query.Create(&newApplication).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateApplication(ctx context.Context, req *domain.Application) (*domain.Application, error) {
	query := r.db.Conn

	uuid := uuid.MustParse(req.ID)
	var existingApplicationDTO dto.Application
	err := query.Where("id = ?", uuid).First(&existingApplicationDTO).Error
	if err != nil {
		return req, err
	}

	existingApplication := dto.ConvertApplicationToDomain(&existingApplicationDTO)

	existingApplication.TicketResponseID = req.TicketResponseID
	existingApplication.UserID = req.UserID
	existingApplication.ManagerID = req.ManagerID
	existingApplication.Status = req.Status
	existingApplication.Type = req.Type
	existingApplication.SubType = req.SubType
	existingApplication.Title = req.Title
	existingApplication.Description = req.Description
	existingApplication.FileName = req.FileName

	newApplication := dto.ConvertApplicationFromDomain(existingApplication)

	err = query.Save(&newApplication).Error
	if err != nil {
		return nil, err
	}

	return existingApplication, nil
}

func (r *Repository) GetApplicationByID(ctx context.Context, id string) (*domain.Application, error) {
	query := r.db.Conn

	uuid := uuid.MustParse(id)
	var applicationDTO dto.Application

	err := query.Where("id = ?", uuid).First(&applicationDTO).Error
	if err != nil {
		return nil, err
	}

	return dto.ConvertApplicationToDomain(&applicationDTO), nil
}

func (r *Repository) GetApplicationsByUserID(ctx context.Context, id string) ([]*domain.Application, error) {
	query := r.db.Conn

	uuid := uuid.MustParse(id)
	var applicationsDTO []*dto.Application
	err := query.Where("user_id = ?", uuid).Find(&applicationsDTO).Error
	if err != nil {
		return nil, err
	}

	var applications []*domain.Application
	for _, applicationDTO := range applicationsDTO {
		applications = append(applications, dto.ConvertApplicationToDomain(applicationDTO))
	}

	return applications, nil
}

func (r *Repository) GetApplications(ctx context.Context, inp *application.ApplicationListInput) ([]*domain.Application, error) {
	query := r.db.Conn

	if inp.Search != "" {
		searchTerm := "%" + inp.Search + "%"
		query = query.Where(
			"title LIKE ? OR description LIKE ?",
			searchTerm, searchTerm,
		)
	}
	if len(inp.Statuses) > 0 {
		query = query.Where("status IN (?)", inp.Statuses)
	}
	if len(inp.Types) > 0 {
		query = query.Where("type IN (?)", inp.Types)
	}
	if len(inp.SubTypes) > 0 {
		query = query.Where("sub_type IN (?)", inp.SubTypes)
	}
	if inp.CreatedAtFrom != "" {
		fromTime, err := utils.GetTimeFromString(inp.CreatedAtFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("created_at >= ?", fromTime)
	}
	if inp.CreatedAtTo != "" {
		toTime, err := utils.GetTimeFromString(inp.CreatedAtTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("created_at <= ?", toTime)
	}
	if inp.UpdatedAtFrom != "" {
		fromTime, err := utils.GetTimeFromString(inp.UpdatedAtFrom)
		if err != nil {
			return nil, err
		}
		query = query.Where("updated_at >= ?", fromTime)
	}
	if inp.UpdatedAtTo != "" {
		toTime, err := utils.GetTimeFromString(inp.UpdatedAtTo)
		if err != nil {
			return nil, err
		}
		query = query.Where("updated_at <= ?", toTime)
	}

	var applicationsDTO []*dto.Application
	err := query.Find(&applicationsDTO).Error
	if err != nil {
		return nil, err
	}

	var applications []*domain.Application
	for _, applicationDTO := range applicationsDTO {
		applications = append(applications, dto.ConvertApplicationToDomain(applicationDTO))
	}

	return applications, nil
}
