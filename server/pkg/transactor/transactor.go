package transactor

import "context"

// Manager - transaction manager that wraps user specified Handler in transaction
type Manager interface {
	InTransaction(ctx context.Context, f Handler) error
}

// Handler - function that wrap atomic actions
type Handler func(ctx context.Context) error
