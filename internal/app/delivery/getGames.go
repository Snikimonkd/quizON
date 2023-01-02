package delivery

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/apiModels"
	"quizON/internal/model/postgres/public/model"
)

type GetGamesUsecase interface {
	GetGames(ctx context.Context, page int32, perPage int32) ([]model.Games, error)
}

func (d *delivery) GetGames(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req apiModels.GetGamesRequest
	err := UnmarshalRequest(r.Body, &req)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	err = apiModels.Validate(req)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	pgGames, err := d.getGamesUsecase.GetGames(ctx, req.Page, req.PerPage)
	if err != nil {
		helpers.HandleError(w, err)
	}

	httpGames := make([]apiModels.Game, 0, len(pgGames))
	for _, v := range pgGames {
		httpGames = append(httpGames, pgToHttpGame(v))
	}

	resp := apiModels.GetGamesResponse{
		Games: httpGames,
	}

	helpers.ResponseWithJson(w, http.StatusOK, resp)
}
