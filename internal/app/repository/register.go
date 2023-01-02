package repository

import (
	"context"
	"fmt"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/postgres/public/model"
	"quizON/internal/model/postgres/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v4"
)

func (r *repository) SetRegisteredTeams(ctx context.Context, gameID int32) (int32, error) {
	stmt := table.Games.UPDATE(table.Games.RegisteredTeams).
		SET(table.Games.RegisteredTeams.ADD(postgres.Int(1))).
		WHERE(table.Games.ID.EQ(postgres.Int32(gameID))).
		RETURNING(table.Games.RegisteredTeams)

	query, args := stmt.Sql()
	var res int32

	err := r.db.QueryRow(ctx, query, args...).Scan(&res)
	if err == pgx.ErrNoRows {
		return -1, helpers.NewHttpError(http.StatusBadRequest, fmt.Errorf("can't update game with id: %w", err), "неправильно указан id игры")
	}
	if err != nil {
		return -1, helpers.NewInternalError(fmt.Errorf("can't set registered teams: %w", err))
	}

	return res, nil
}

func (r *repository) Register(ctx context.Context, registration model.Registrations) error {
	stmt := table.Registrations.INSERT(
		table.Registrations.GameID,
		table.Registrations.TeamName,
		table.Registrations.CaptainName,
		table.Registrations.Phone,
		table.Registrations.Telega,
		table.Registrations.Amount,
		table.Registrations.RegistrationNumber,
	).VALUES(
		registration.GameID,
		registration.TeamName,
		registration.CaptainName,
		registration.Phone,
		registration.Telega,
		registration.Amount,
		registration.RegistrationNumber,
	)

	query, args := stmt.Sql()

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return helpers.NewInternalError(fmt.Errorf("can't insert registration: %w", err))
	}

	return nil
}
