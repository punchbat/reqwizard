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

func (uc *UseCase) CreateApplication(ctx context.Context, inp *application.CreateApplicationInput) error {
	userEntity, err := uc.authRepo.GetUserByID(ctx, inp.ID)
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
	if inp.File != nil {
		fileExt := filepath.Ext(inp.FileName)
		newFileName := uuid.New().String() + fileExt
		filePath := "uploads/applications/" + newFileName

		fileBytes, err := ioutil.ReadAll(inp.File)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filePath, fileBytes, 0644)
		if err != nil {
			return err
		}

		entity.FileName = newFileName
	}

	err = uc.repo.CreateApplication(ctx, entity)
	if err != nil {
		return err
	}

	userEntity.ApplicationCreatedAt = time.Now()

	err = uc.authRepo.UpdateUser(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) GetFile(ctx context.Context, userID string, fileName string) ([]byte, string, error) {
	userEntity, err := uc.authRepo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, "", err
	}

	applicationEntities, err := uc.repo.GetApplicationsByUserID(ctx, &application.ApplicationListInput{
		ID: userEntity.ID,
	})
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
		return nil, "", errors.New("User doesn`t have access to the file")
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

func (uc *UseCase) GetApplicationByID(ctx context.Context, id string) (*domain.Application, error) {
	application, err := uc.repo.GetApplicationByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return application, nil
}

func (uc *UseCase) GetApplicationsByUserID(ctx context.Context, inp *application.ApplicationListInput) ([]*domain.Application, error) {
	applications, err := uc.repo.GetApplicationsByUserID(ctx, inp)
	if err != nil {
		return nil, err
	}

	return applications, nil
}

// * manager.
func (uc *UseCase) GetApplications(ctx context.Context, inp *application.ApplicationListInput) ([]*domain.Application, error) {
	applications, err := uc.repo.GetApplications(ctx, inp)
	if err != nil {
		return nil, err
	}

	return applications, nil
}