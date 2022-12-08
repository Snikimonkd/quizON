package usecase

import (
	"context"
	"quizON/internal/app/repository"
	"quizON/internal/model/postgres/public/model"
)

// CreateGameRepository - интерфейс репозитория для юскейса CreateGame
type CreateGameRepository interface {
	CreateGame(ctx context.Context, game model.Games) (int32, error)
}

// createGameUsecase - реализация юскейса CreateGame
type createGameUsecase struct {
	createGameRepository CreateGameRepository
	commonRepository     repository.CommonRepository
}

// NewCreateGameUsecase - конструктор createGameUsecase
func NewCreateGameUsecase(createGameRepository CreateGameRepository, commonRepository repository.CommonRepository) *createGameUsecase {
	return &createGameUsecase{
		createGameRepository: createGameRepository,
		commonRepository:     commonRepository,
	}
}

func (c *createGameUsecase) CreateGame(ctx context.Context, game model.Games) (model.Games, error) {
	var err error
	game.ID, err = c.createGameRepository.CreateGame(ctx, game)
	if err != nil {
		return model.Games{}, err
	}

	return game, nil
}
