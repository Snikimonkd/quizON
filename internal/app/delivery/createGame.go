package delivery

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/apiModels"
	"quizON/internal/model/postgres/public/model"
)

type CreateGameUsecase interface {
	CreateGame(ctx context.Context, game model.Games) (model.Games, error)
}

func (d *delivery) CreateGame(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req apiModels.CreateGameRequest
	err := MarshalRequest(r.Body, &req)
	if err != nil {
		helpers.HandleHttpError(w, err)
		return
	}

	err = apiModels.Validate(req)
	if err != nil {
		helpers.HandleHttpError(w, err)
		return
	}

	pgGame, err := httpToPgGame(ctx, req)
	if err != nil {
		helpers.HandleHttpError(w, err)
		return
	}

	pgGame, err = d.createGameUsecase.CreateGame(ctx, pgGame)
	if err != nil {
		helpers.HandleHttpError(w, err)
		return
	}

	resp := pgToHttpGame(pgGame)

	helpers.ResponseWithJson(w, http.StatusOK, resp)
}
