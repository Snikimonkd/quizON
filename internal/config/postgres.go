package config

import (
	"context"
	"quizON/internal/logger"

	"github.com/jackc/pgx/v4"
)

// ConnectToPostgres - подключается к postgres
func ConnectToPostgres(ctx context.Context) *pgx.Conn {
	db, err := pgx.Connect(ctx, GlobalConfig.Database.DSN)
	if err != nil {
		logger.Fatalf("Can't connect to db: %v\n", err)
	}

	return db
}
