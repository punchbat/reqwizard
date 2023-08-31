package repository

import (
	"context"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/ticketResponse"
	"reqwizard/internal/routes/ticketResponse/repository/dto"
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

func (r Repository) CreateTicketResponse(ctx context.Context, req *domain.TicketResponse) error {
	query := r.db.Conn

	newTicketResponse := dto.ConvertTicketResponseFromDomain(req)

	err := query.Create(&newTicketResponse).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetTicketResponsesByManagerID(ctx context.Context, inp *ticketResponse.TicketResponseListInput) ([]*domain.TicketResponse, error) {
	query := r.db.Conn

	uuid := uuid.MustParse(inp.ID)
	query = query.Where("manager_id = ?", uuid)

	if inp.Search != "" {
		searchTerm := "%" + inp.Search + "%"
		query = query.Where(
			"text LIKE ?",
			searchTerm,
		)
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

	var ticketResponsesDTO []*dto.TicketResponse
	err := query.Find(&ticketResponsesDTO).Error
	if err != nil {
		return nil, err
	}

	var ticketResponses []*domain.TicketResponse
	for _, ticketResponseDTO := range ticketResponsesDTO {
		ticketResponses = append(ticketResponses, dto.ConvertTicketResponseToDomain(ticketResponseDTO))
	}

	return ticketResponses, nil
}

func (r *Repository) GetTicketResponseByID(ctx context.Context, id string) (*domain.TicketResponse, error) {
	query := r.db.Conn

	uuid := uuid.MustParse(id)
	var ticketResponseDTO dto.TicketResponse

	err := query.Where("id = ?", uuid).First(&ticketResponseDTO).Error
	if err != nil {
		return nil, err
	}

	return dto.ConvertTicketResponseToDomain(&ticketResponseDTO), nil
}

func (r *Repository) GetTicketResponsesByUserID(ctx context.Context, inp *ticketResponse.TicketResponseListInput) ([]*domain.TicketResponse, error) {
	query := r.db.Conn

	uuid := uuid.MustParse(inp.ID)
	query = query.Where("user_id = ?", uuid)

	if inp.Search != "" {
		searchTerm := "%" + inp.Search + "%"
		query = query.Where(
			"text LIKE ?",
			searchTerm,
		)
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

	var ticketResponsesDTO []*dto.TicketResponse
	err := query.Find(&ticketResponsesDTO).Error
	if err != nil {
		return nil, err
	}

	var ticketResponses []*domain.TicketResponse
	for _, ticketResponseDTO := range ticketResponsesDTO {
		ticketResponses = append(ticketResponses, dto.ConvertTicketResponseToDomain(ticketResponseDTO))
	}

	return ticketResponses, nil
}

func (r *Repository) GetTicketResponses(ctx context.Context) ([]*domain.TicketResponse, error) {
	query := r.db.Conn

	var ticketResponsesDTO []*dto.TicketResponse
	err := query.Find(&ticketResponsesDTO).Error
	if err != nil {
		return nil, err
	}

	var ticketResponses []*domain.TicketResponse
	for _, ticketResponseDTO := range ticketResponsesDTO {
		ticketResponses = append(ticketResponses, dto.ConvertTicketResponseToDomain(ticketResponseDTO))
	}

	return ticketResponses, nil
}