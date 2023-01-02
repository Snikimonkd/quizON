package repository

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgconn"
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

// Executor - интерфейс для работы с базой
type Executor interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
}

func selectx(ctx context.Context, db Executor, dest interface{}, stmt postgres.Statement) error {
	query, args := stmt.Sql()

	return pgxscan.Select(ctx, db, dest, query, args...)
}

func getx(ctx context.Context, db Executor, dest interface{}, stmt postgres.Statement) error {
	query, args := stmt.Sql()

	return pgxscan.Get(ctx, db, dest, query, args...)
}

func execx(ctx context.Context, ex Executor, stmt postgres.Statement) (commandTag pgconn.CommandTag, err error) {
	query, args := stmt.Sql()

	return ex.Exec(ctx, query, args...)
}
