package delivery

import (
	"context"
	"encoding/json"
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
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		retErr := helpers.NewHttpError(http.StatusBadRequest, err, err.Error())
		helpers.HandleHttpError(w, retErr)
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
