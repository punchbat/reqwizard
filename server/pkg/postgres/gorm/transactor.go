package gorm

import (
	"context"
	"reqwizard/pkg/transactor"

	"gorm.io/gorm"
)

type TransactionManager struct {
	Gorm *Gorm
}

func NewTransactionManager(gorm *Gorm) *TransactionManager {
	return &TransactionManager{
		Gorm: gorm,
	}
}

// TODO: fix
func (t *TransactionManager) InTransaction(ctx context.Context, f transactor.Handler) error {
	return t.Gorm.Conn.Transaction(func(tx *gorm.DB) error {
		return f(ctx)
	})
}
