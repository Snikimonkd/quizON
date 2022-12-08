package repository

import (
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

var NotFoundError = errors.New("not found")

// repository - слой хранения
type repository struct {
	db *pgx.Conn
}

// NewRepository - конструктор для слоя доставки
func NewRepository(db *pgx.Conn) *repository {
	return &repository{
		db: db,
	}
}
