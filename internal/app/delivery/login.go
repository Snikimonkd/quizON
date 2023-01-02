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
	err := UnmarshalRequest(r.Body, &user)
	if err != nil {
		helpers.HandleError(w, helpers.NewHttpError(http.StatusBadRequest, err, "can't unmarshal body"))
		return
	}

	err = apiModels.Validate(user)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	cookie, err := d.loginUsecase.Login(ctx, user)
	if err != nil {
		helpers.HandleError(w, err)
		return
	}

	SetCookie(w, cookie)
	helpers.ResponseWithJson(w, http.StatusOK, nil)
}
