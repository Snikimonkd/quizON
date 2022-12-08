package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/app/repository"
	"quizON/internal/app/usecase"

	"github.com/jackc/pgx/v4"
)

// delivery - слой доставки
type delivery struct {
	loginUsecase      LoginUsecase
	createGameUsecase CreateGameUsecase
}

// NewDelivery - конструктор для слоя доставки
func NewDelivery(db *pgx.Conn) *delivery {
	commonRepository := repository.NewCommonRepository(db)

	repo := repository.NewRepository(db)

	loginUsecase := usecase.NewLoginUsecase(repo, commonRepository)
	createGameUsecase := usecase.NewCreateGameUsecase(repo, commonRepository)

	return &delivery{
		loginUsecase:      loginUsecase,
		createGameUsecase: createGameUsecase,
	}
}

func MarshalRequest[T any](body io.ReadCloser, value *T) error {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(value)
	if err != nil {
		return helpers.NewHttpError(http.StatusBadRequest, err, helpers.BadRequest)
	}

	return nil
}
