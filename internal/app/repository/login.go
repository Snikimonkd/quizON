package repository

import (
	"context"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/postgres/public/model"
	"quizON/internal/model/postgres/public/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *repository) GetUserByLogin(ctx context.Context, tx pgx.Tx, login string) (model.Users, error) {
	stmt := table.Users.SELECT(
		table.Users.ID,
		table.Users.Login,
		table.Users.Password,
	).WHERE(table.Users.Login.EQ(postgres.String(login)))

	query, args := stmt.Sql()

	var user model.Users

	err := tx.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Login, &user.Password)
	if errors.Is(err, pgx.ErrNoRows) {
		return model.Users{}, helpers.NewHttpError(http.StatusForbidden, err, helpers.WrongLoginOrPassword)
	}
	if err != nil {
		return model.Users{}, helpers.NewHttpError(http.StatusInternalServerError, err, helpers.EmptyResponse)
	}

	return user, nil
}

func (r *repository) CreateCookie(ctx context.Context, tx pgx.Tx, id int32) (model.Cookies, error) {
	cookie := model.Cookies{
		Value:  uuid.New(),
		UserID: id,
	}

	stmt := table.Cookies.INSERT(
		table.Cookies.UserID,
		table.Cookies.Value,
	).VALUES(
		id,
		uuid.New(),
	).RETURNING(
		table.Cookies.ExpiresAt,
	)
	query, args := stmt.Sql()

	err := tx.QueryRow(ctx, query, args...).Scan(&cookie.ExpiresAt)
	if err != nil {
		return model.Cookies{}, helpers.NewHttpError(http.StatusInternalServerError, err, helpers.EmptyResponse)
	}

	return cookie, nil
}
