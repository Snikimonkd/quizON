package repository

import (
	"context"
	"fmt"
	"quizON/internal/app/helpers"
	"quizON/internal/model/postgres/public/model"
	"quizON/internal/model/postgres/public/table"
)

func (r *repository) GetGames(ctx context.Context, page int32, perPage int32) ([]model.Games, error) {
	stmt := table.Games.SELECT(
		table.Games.ID,
		table.Games.Name,
		table.Games.Description,
		table.Games.Date,
		table.Games.TeamsAmount,
		table.Games.PricePerPerson,
		table.Games.Location,
		table.Games.CreatedBy,
		table.Games.RegisteredTeams,
	).ORDER_BY(table.Games.Date.DESC()).
		LIMIT(int64(perPage)).
		OFFSET(int64((page - 1) * perPage))

	query, args := stmt.Sql()
	var games []model.Games

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, helpers.NewInternalError(fmt.Errorf("can't select games: %w", err))
	}
	defer rows.Close()
	for rows.Next() {
		var game model.Games
		err = rows.Scan(
			&game.ID,
			&game.Name,
			&game.Description,
			&game.Date,
			&game.TeamsAmount,
			&game.PricePerPerson,
			&game.Location,
			&game.CreatedBy,
			&game.RegisteredTeams,
		)
		if err != nil {
			return nil, helpers.NewInternalError(fmt.Errorf("can't scan games: %w", err))
		}
		games = append(games, game)
	}

	return games, nil
}
