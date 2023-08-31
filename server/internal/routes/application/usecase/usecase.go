package usecase

import (
	"context"
	"errors"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"reqwizard/internal/domain"
	"time"

	"reqwizard/internal/routes/application"
	"reqwizard/internal/routes/auth"

	"github.com/google/uuid"
)

type UseCase struct {
	repo     application.Repository
	authRepo auth.Repository
}

func NewUseCase(repo application.Repository, authRepo auth.Repository) *UseCase {
	return &UseCase{
		repo:     repo,
		authRepo: authRepo,
	}
}

func (a *UseCase) CreateApplication(ctx context.Context, inp *application.CreateApplicationInput) error {
	userEntity, err := a.authRepo.GetUserByID(ctx, inp.ID)
	if err != nil {
		return err
	}

	// // * Проверяем, была ли уже создана заявка пользователем
	// now := time.Now().UTC()
	// // * today := time.Now().UTC().Truncate(24 * time.Hour) // Обнуляем время
	// today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	// if !userEntity.ApplicationCreatedAt.IsZero() && userEntity.ApplicationCreatedAt.After(today) {
	// 	return errors.New("user can create only 1 application per day")
	// }

	// * Создаем Заявку со статус Ожидание
	entity := &domain.Application{
		ID:          uuid.New().String(),
		UserID:      inp.ID,
		Status:      domain.ApplicationStatusWaiting,
		Type:        domain.ApplicationType(inp.Type),
		SubType:     domain.ApplicationSubType(inp.SubType),
		Title:       inp.Title,
		Description: inp.Description,
	}

	// * Проверяем наличие файла
	if len(inp.File) > 0 {
		fileName := uuid.New().String() + inp.FileExtension
		filePath := "uploads/applications/" + fileName

		err = ioutil.WriteFile(filePath, inp.File, 0644)
		if err != nil {
			return err
		}

		entity.FileName = fileName
	}

	err = a.repo.CreateApplication(ctx, entity)
	if err != nil {
		return err
	}

	userEntity.ApplicationCreatedAt = time.Now()

	err = a.authRepo.UpdateUser(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (a *UseCase) GetFile(ctx context.Context, userID string, fileName string) ([]byte, string, error) {
	userEntity, err := a.authRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, "", err
	}

	applicationEntities, err := a.repo.GetApplicationsByUserID(ctx, userEntity.ID)
	if err != nil {
		return nil, "", err
	}

	hasFileName := false
	for i := 0; i < len(applicationEntities); i++ {
		if applicationEntities[i].FileName == fileName {
			hasFileName = true
		}
	}

	if !hasFileName {
		return nil, "", errors.New("User doesnt have access to the file")
	}

	filePath := "uploads/applications/" + fileName
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, "", err
	}

	mimeType := mime.TypeByExtension(filepath.Ext(filePath))

	return fileContents, mimeType, nil
}

func (a *UseCase) GetApplicationByID(ctx context.Context, id string) (*domain.Application, error) {
	application, err := a.repo.GetApplicationByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return application, nil
}

func (a *UseCase) GetApplicationsByUserID(ctx context.Context, id string) ([]*domain.Application, error) {
	applications, err := a.repo.GetApplicationsByUserID(ctx, id)

	if err != nil {
		return nil, err
	}

	return applications, nil
}

// * manager.
func (a *UseCase) GetApplications(ctx context.Context, inp *application.ApplicationListInput) ([]*domain.Application, error) {
	applications, err := a.repo.GetApplications(ctx, inp)
	if err != nil {
		return nil, err
	}

	return applications, nil
}