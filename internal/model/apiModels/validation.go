package apiModels

import (
	"net/http"
	"quizON/internal/app/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(in interface{}) error {
	err := validate.Struct(in)
	if err != nil {
		return helpers.NewHttpError(http.StatusBadRequest, errors.Wrap(err, "validation error"), helpers.ValidationError)
	}

	return nil
}
