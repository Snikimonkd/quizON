package usecase

import (
	"context"
	"quizON/internal/model/postgres/public/model"
)

// GetGamesRepository - интерфейс репозитория для юскейса GetGames
type GetGamesRepository interface {
	GetGames(ctx context.Context, page int32, perPage int32) ([]model.Games, error)
}

// getGamesUsecase - реализация юскейса GetGames
type getGamesUsecase struct {
	getGamesRepository GetGamesRepository
}

// NewGetGamesUsecase - конструктор getGamesUsecase
func NewGetGamesUsecase(getGamesRepository GetGamesRepository) *getGamesUsecase {
	return &getGamesUsecase{
		getGamesRepository: getGamesRepository,
	}
}

func (c *getGamesUsecase) GetGames(ctx context.Context, page int32, perPage int32) ([]model.Games, error) {
	games, err := c.getGamesRepository.GetGames(ctx, page, perPage)
	if err != nil {
		return []model.Games{}, err
	}

	return games, nil
}
