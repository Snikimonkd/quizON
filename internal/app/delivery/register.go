package delivery

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/apiModels"
	"quizON/internal/model/postgres/public/model"
)

type RegisterUsecase interface {
	Register(ctx context.Context, registerReq model.Registrations) (int32, error)
}

func (d *delivery) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req apiModels.RegisterRequest
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

	registration := httpToPgRegister(req)

	num, err := d.registerUsecase.Register(ctx, registration)
	if err != nil {
		helpers.HandleError(w, err)
	}

	response := apiModels.RegisterResponse{
		RegistrationNumber: num,
	}

	helpers.ResponseWithJson(w, http.StatusOK, response)
}
