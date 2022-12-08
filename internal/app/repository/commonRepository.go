package repository

import (
	"context"
	"fmt"
	"quizON/internal/app/helpers"
	"quizON/internal/logger"

	"github.com/jackc/pgx/v4"
)

// CommonRepository - общий репозиторий
type CommonRepository interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
	CommitTx(ctx context.Context, tx pgx.Tx) error
	RollBackUnlessCommitted(ctx context.Context, tx pgx.Tx)
}

// commonRepository - реализация общего репозитория
type commonRepository struct {
	db *pgx.Conn
}

// NewCommonRepository возвращает новый commonRepository
func NewCommonRepository(db *pgx.Conn) *commonRepository {
	return &commonRepository{db: db}
}

// BeginTx - начать транзакцию
func (c *commonRepository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := c.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, helpers.NewInternalError(fmt.Errorf("can't begin transaction: %w", err))
	}

	return tx, nil
}

// CommitTx - закоммитить транзакцию
func (c *commonRepository) CommitTx(ctx context.Context, tx pgx.Tx) error {
	err := tx.Commit(ctx)
	if err != nil {
		return helpers.NewInternalError(fmt.Errorf("can't cpmmit transaction: %w", err))
	}

	return nil
}

// RollBackUnlessCommitted - роллбэк, если транзакция не закоммичена
func (c *commonRepository) RollBackUnlessCommitted(ctx context.Context, tx pgx.Tx) {
	if tx == nil {
		return
	}

	err := tx.Rollback(ctx)
	if err == pgx.ErrTxClosed {
		return
	}

	if err != nil {
		logger.Errorf("can't rollback transaction: %v", err)
	}
}
