package usecase

import (
	"context"
	"quizON/internal/app/repository"
	"quizON/internal/model/postgres/public/model"
)

// RegisterRepository - интерфейс репозитория для юскейса Register
type RegisterRepository interface {
	Register(ctx context.Context, registration model.Registrations) error
	SetRegisteredTeams(ctx context.Context, gameID int32) (int32, error)
}

// registerUsecase - реализация юскейса Register
type registerUsecase struct {
	registerRepository RegisterRepository
	commonRepository   repository.CommonRepository
}

// NewRegisterUsecase - конструктор registerUsecase
func NewRegisterUsecase(registerRepository RegisterRepository, commonRepository repository.CommonRepository) *registerUsecase {
	return &registerUsecase{
		registerRepository: registerRepository,
		commonRepository:   commonRepository,
	}
}

func (r *registerUsecase) Register(ctx context.Context, registration model.Registrations) (int32, error) {
	tx, err := r.commonRepository.BeginTx(ctx)
	if err != nil {
		return -1, err
	}
	defer r.commonRepository.RollBackUnlessCommitted(ctx, tx)

	registration.RegistrationNumber, err = r.registerRepository.SetRegisteredTeams(ctx, registration.GameID)
	if err != nil {
		return -1, err
	}

	err = r.registerRepository.Register(ctx, registration)
	if err != nil {
		return -1, err
	}

	err = r.commonRepository.CommitTx(ctx, tx)
	if err != nil {
		return -1, err
	}

	return registration.RegistrationNumber, nil
}
