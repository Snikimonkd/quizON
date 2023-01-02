package delivery

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/apiModels"
	"quizON/internal/model/postgres/public/model"
)

type GetRegistrationsUsecase interface {
	GetRegistrations(ctx context.Context, gameID int32) ([]model.Registrations, error)
}

func (d *delivery) GetRegistrations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req apiModels.GetRegistrationsRequest
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

	registrations, err := d.getRegistrationsUsecase.GetRegistrations(ctx, req.GameID)
	if err != nil {
		helpers.HandleError(w, err)
	}

	responseRegistrations := make([]apiModels.Registration, 0, len(registrations))
	for _, v := range registrations {
		responseRegistrations = append(responseRegistrations, pgToHttpRegistration(v))
	}

	response := apiModels.GetRegistrationsResponse{
		Registrations: responseRegistrations,
	}

	helpers.ResponseWithJson(w, http.StatusOK, response)
}
