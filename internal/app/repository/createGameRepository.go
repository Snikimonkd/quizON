package repository

import (
	"context"
	"fmt"
	"quizON/internal/app/helpers"
	"quizON/internal/model/postgres/public/model"
	"quizON/internal/model/postgres/public/table"
)

func (r *repository) CreateGame(ctx context.Context, game model.Games) (int32, error) {
	stmt := table.Games.INSERT(
		table.Games.Name,
		table.Games.Description,
		table.Games.Date,
		table.Games.TeamsAmount,
		table.Games.PricePerPerson,
		table.Games.Location,
		table.Games.CreatedBy,
	).VALUES(
		game.Name,
		game.Description,
		game.Date,
		game.TeamsAmount,
		game.PricePerPerson,
		game.Location,
		game.CreatedBy,
	).RETURNING(table.Games.ID)
	query, args := stmt.Sql()

	var id int32
	err := r.db.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return -1, helpers.NewInternalError(fmt.Errorf("can't create game: %w", err))
	}

	return id, nil
}
