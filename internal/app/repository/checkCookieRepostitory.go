package repository

import (
	"context"
	"errors"
	"fmt"
	"quizON/internal/app/helpers"
	"quizON/internal/model/postgres/public/table"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

// CheckCookieRepository - интерфейс для проверки куки
type CheckCookieRepository interface {
	CheckCookie(ctx context.Context, value uuid.UUID) (int32, error)
}

// checkCookieRepository - реализация интерфейса проверки куки
type checkCookieRepository struct {
	db *pgx.Conn
}

// NewCheckCookieRepository - конструктор для репозитория
func NewCheckCookieRepository(db *pgx.Conn) *checkCookieRepository {
	return &checkCookieRepository{
		db: db,
	}
}

// CheckCookie - проверить куку
func (c *checkCookieRepository) CheckCookie(ctx context.Context, value uuid.UUID) (int32, error) {
	stmt := table.Cookies.SELECT(table.Cookies.UserID).
		WHERE(table.Cookies.Value.EQ(postgres.UUID(value)).
			AND(table.Cookies.ExpiresAt.GT(postgres.TimestampzT(time.Now()))))
	query, args := stmt.Sql()

	var id int32
	err := c.db.QueryRow(ctx, query, args...).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return -1, NotFoundError
	}
	if err != nil {
		return -1, helpers.NewInternalError(fmt.Errorf("can't check cookie: %w", err))
	}

	return id, nil
}
