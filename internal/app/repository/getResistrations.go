package repository

import (
	"context"
	"fmt"
	"quizON/internal/app/helpers"
	"quizON/internal/model/postgres/public/model"
	"quizON/internal/model/postgres/public/table"
)

func (r *repository) GetRegistrations(ctx context.Context, gameID int32) ([]model.Registrations, error) {
	stmt := table.Registrations.SELECT(
		table.Registrations.TeamName,
		table.Registrations.CaptainName,
		table.Registrations.Phone,
		table.Registrations.Telega,
		table.Registrations.Amount,
		table.Registrations.RegistrationNumber,
		table.Registrations.CreatedAt,
	).ORDER_BY(table.Registrations.RegistrationNumber.ASC())

	query, args := stmt.Sql()
	var registrations []model.Registrations

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, helpers.NewInternalError(fmt.Errorf("can't select games: %w", err))
	}
	defer rows.Close()
	for rows.Next() {
		var registration model.Registrations
		err = rows.Scan(
			&registration.TeamName,
			&registration.CaptainName,
			&registration.Phone,
			&registration.Telega,
			&registration.Amount,
			&registration.RegistrationNumber,
			&registration.CreatedAt,
		)
		if err != nil {
			return nil, helpers.NewInternalError(fmt.Errorf("can't scan games: %w", err))
		}
		registrations = append(registrations, registration)
	}

	return registrations, nil
}
