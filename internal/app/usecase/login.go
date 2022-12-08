package usecase

import (
	"context"
	"fmt"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/app/repository"
	"quizON/internal/model/apiModels"
	"quizON/internal/model/postgres/public/model"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// LoginRepository - интерфейс репозитория для юскейса Login
type LoginRepository interface {
	GetUserByLogin(ctx context.Context, tx pgx.Tx, login string) (model.Users, error)
	CreateCookie(ctx context.Context, tx pgx.Tx, id int32) (model.Cookies, error)
}

// loginUsecase - реализация юскейса Login
type loginUsecase struct {
	loginRepository  LoginRepository
	commonRepository repository.CommonRepository
}

// NewLoginUsecase - конструктор loginUsecase
func NewLoginUsecase(loginRepository LoginRepository, commonRepository repository.CommonRepository) *loginUsecase {
	return &loginUsecase{
		loginRepository:  loginRepository,
		commonRepository: commonRepository,
	}
}

// Login - юскейс Login
func (l *loginUsecase) Login(ctx context.Context, loginRequest apiModels.LoginRequest) (model.Cookies, error) {
	tx, err := l.commonRepository.BeginTx(ctx)
	if err != nil {
		return model.Cookies{}, err
	}
	defer l.commonRepository.RollBackUnlessCommitted(ctx, tx)

	user, err := l.loginRepository.GetUserByLogin(ctx, tx, loginRequest.Login)
	if err != nil {
		return model.Cookies{}, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(loginRequest.Password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return model.Cookies{}, helpers.NewHttpError(http.StatusForbidden, err, helpers.AuthenticationError)
	}
	if err != nil {
		return model.Cookies{}, helpers.NewInternalError(fmt.Errorf("can't comapre passwords: %w", err))
	}

	cookie, err := l.loginRepository.CreateCookie(ctx, tx, user.ID)
	if err != nil {
		return model.Cookies{}, err
	}

	err = l.commonRepository.CommitTx(ctx, tx)
	if err != nil {
		return model.Cookies{}, helpers.NewHttpError(http.StatusInternalServerError, errors.Wrap(err, "can't commit transaction"), helpers.Empty)
	}

	return cookie, nil
}
