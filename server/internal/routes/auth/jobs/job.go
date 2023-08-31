package jobs

import (
	"context"
	"log"
	"reqwizard/internal/routes/auth"

	"github.com/robfig/cron/v3"
)

type AuthJobScheduler struct {
	useCase auth.UseCase
}

func NewAuthJobScheduler(useCase auth.UseCase) *AuthJobScheduler {
	return &AuthJobScheduler{
		useCase: useCase,
	}
}

func (uc *AuthJobScheduler) Start(c *cron.Cron) {
	go func() {
		_, err := c.AddFunc("0 0 * * *", uc.RemoveUnverifiedUsers) // Ежедневно в полночь UTC
		if err != nil {
			log.Fatalf("failed to start job RemoveUnverifiedUsers %v", err)
		}
	}()
}

func (uc *AuthJobScheduler) RemoveUnverifiedUsers() {
	log.Printf("started job creating new partition")

	err := uc.useCase.RemoveUnverifiedUsers(context.Background())
	if err != nil {
		log.Fatalf("failed to create new partition for dlh table: %v", err)
	}

	log.Printf("finished job creating new partition")
}