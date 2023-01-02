package delivery

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/app/middleware"
	"quizON/internal/model/apiModels"
	"quizON/internal/model/postgres/public/model"

	"github.com/pkg/errors"
)

func httpToPgGame(ctx context.Context, in apiModels.CreateGameRequest) (model.Games, error) {
	ctxID := ctx.Value(middleware.CtxUserID)
	if ctxID == nil {
		return model.Games{}, helpers.NewHttpError(http.StatusUnauthorized, errors.New("can't find id in ctx"), helpers.UnauthorizedError)
	}

	createdBy, ok := ctxID.(int32)
	if !ok {
		return model.Games{}, helpers.NewHttpError(http.StatusUnauthorized, errors.New("can't cast cookie id to int32"), helpers.UnauthorizedError)
	}

	return model.Games{
		Name:           in.Name,
		Description:    in.Description,
		Date:           in.Date,
		TeamsAmount:    in.TeamsAmount,
		PricePerPerson: in.PricePerPerson,
		Location:       in.Location,
		CreatedBy:      createdBy,
	}, nil
}

func pgToHttpGame(in model.Games) apiModels.Game {
	return apiModels.Game{
		ID:              in.ID,
		Name:            in.Name,
		Description:     in.Description,
		Date:            in.Date,
		TeamsAmount:     in.TeamsAmount,
		RegisteredTeams: in.RegisteredTeams,
		PricePerPerson:  in.PricePerPerson,
		Location:        in.Location,
	}
}

func httpToPgRegistration(in apiModels.RegisterRequest) model.Registrations {
	return model.Registrations{
		GameID:      in.GameID,
		TeamName:    in.TeamName,
		CaptainName: in.CaptainName,
		Phone:       in.Phone,
		Telega:      in.Telega,
		Amount:      in.Amount,
	}
}

func pgToHttpRegistration(in model.Registrations) apiModels.Registration {
	return apiModels.Registration{
		TeamName:           in.TeamName,
		CaptainName:        in.CaptainName,
		Phone:              in.Phone,
		Telega:             in.Telega,
		Amount:             in.Amount,
		RegistrationNumber: in.RegistrationNumber,
		CreatedAt:          in.CreatedAt,
	}
}
