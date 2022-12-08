package delivery

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/apiModels"
	"quizON/internal/model/postgres/public/model"
)

type LoginUsecase interface {
	Login(ctx context.Context, loginRequest apiModels.LoginRequest) (model.Cookies, error)
}

func (d *delivery) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var user apiModels.LoginRequest
	err := MarshalRequest(r.Body, &user)
	if err != nil {
		helpers.HandleHttpError(w, err)
		return
	}

	err = apiModels.Validate(user)
	if err != nil {
		helpers.HandleHttpError(w, err)
		return
	}

	cookie, err := d.loginUsecase.Login(ctx, user)
	if err != nil {
		helpers.HandleHttpError(w, err)
		return
	}

	SetCookie(w, cookie)
	helpers.ResponseWithJson(w, http.StatusOK, nil)
}
