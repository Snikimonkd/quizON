package repository

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/logger"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

// commonRepository - репозиторий с общими функциями
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
		return nil, helpers.NewHttpError(http.StatusInternalServerError, errors.Wrap(err, "can't begin transaction"), helpers.EmptyResponse)
	}

	return tx, nil
}

// CommitTx - закоммитить транзакцию
func (c *commonRepository) CommitTx(ctx context.Context, tx pgx.Tx) error {
	err := tx.Commit(ctx)
	if err != nil {
		return helpers.NewHttpError(http.StatusInternalServerError, errors.Wrap(err, "can't begin transaction"), helpers.EmptyResponse)
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
