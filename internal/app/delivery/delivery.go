package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/app/repository"
	"quizON/internal/app/usecase"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

// delivery - слой доставки
type delivery struct {
	loginUsecase      LoginUsecase
	createGameUsecase CreateGameUsecase
	getGamesUsecase   GetGamesUsecase
	registerUsecase   RegisterUsecase
}

// NewDelivery - конструктор для слоя доставки
func NewDelivery(db *pgx.Conn) *delivery {
	commonRepository := repository.NewCommonRepository(db)
	repository := repository.NewRepository(db)

	loginUsecase := usecase.NewLoginUsecase(repository, commonRepository)
	createGameUsecase := usecase.NewCreateGameUsecase(repository, commonRepository)
	getGamesUsecase := usecase.NewGetGamesUsecase(repository)
	registerUsecase := usecase.NewRegisterUsecase(repository, commonRepository)

	return &delivery{
		loginUsecase:      loginUsecase,
		createGameUsecase: createGameUsecase,
		getGamesUsecase:   getGamesUsecase,
		registerUsecase:   registerUsecase,
	}
}

func UnmarshalRequest[T any](body io.ReadCloser, value *T) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(value)
	if err != nil {
		return helpers.NewHttpError(http.StatusBadRequest, errors.Wrap(err, "unmarshal error"), helpers.BadRequest)
	}

	return nil
}
