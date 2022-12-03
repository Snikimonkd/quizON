package delivery

import (
	"quizON/internal/app/repository"
	"quizON/internal/app/usecase"

	"github.com/jackc/pgx/v4"
)

// delivery - слой доставки
type delivery struct {
	loginUsecase LoginUsecase
}

// NewDelivery - конструктор для слоя доставки
func NewDelivery(db *pgx.Conn) *delivery {
	commonRepository := repository.NewCommonRepository(db)

	loginRepository := repository.NewRepository(db)

	loginUsecase := usecase.NewLoginUsecase(loginRepository, commonRepository)

	return &delivery{
		loginUsecase: loginUsecase,
	}
}
