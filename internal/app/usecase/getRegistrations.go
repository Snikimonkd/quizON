package usecase

import (
	"context"
	"quizON/internal/model/postgres/public/model"
)

// GetRegistrationsRepository - интерфейс репозитория для юскейса GetRegistrations
type GetRegistrationsRepository interface {
	GetRegistrations(ctx context.Context, gameID int32) ([]model.Registrations, error)
}

// getRegistrationsUsecase - реализация юскейса GetRegistrations
type getRegistrationsUsecase struct {
	getRegistrationsRepository GetRegistrationsRepository
}

// NewGetRegistrationsUsecase - конструктор getRegistrationsUsecase
func NewGetRegistrationsUsecase(getRegistrationsRepository GetRegistrationsRepository) *getRegistrationsUsecase {
	return &getRegistrationsUsecase{
		getRegistrationsRepository: getRegistrationsRepository,
	}
}

func (c *getRegistrationsUsecase) GetRegistrations(ctx context.Context, gameID int32) ([]model.Registrations, error) {
	registrations, err := c.getRegistrationsRepository.GetRegistrations(ctx, gameID)
	if err != nil {
		return []model.Registrations{}, err
	}

	return registrations, nil
}
