package repository

import (
	"context"
	"fmt"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/model/postgres/public/model"
	"quizON/internal/model/postgres/public/table"

	"github.com/go-jet/jet/v2/postgres"
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
		return model.Users{}, helpers.NewHttpError(http.StatusForbidden, err, helpers.AuthenticationError)
	}
	if err != nil {
		return model.Users{}, helpers.NewInternalError(fmt.Errorf("can't get user by login: %w", err))
	}

	return user, nil
}

func (r *repository) CreateCookie(ctx context.Context, tx pgx.Tx, id int32) (model.Cookies, error) {
	cookie := model.Cookies{
		UserID: id,
	}

	stmt := table.Cookies.INSERT(
		table.Cookies.UserID,
	).VALUES(
		id,
	).RETURNING(
		table.Cookies.Value,
		table.Cookies.ExpiresAt,
	)
	query, args := stmt.Sql()

	err := tx.QueryRow(ctx, query, args...).Scan(&cookie.Value, &cookie.ExpiresAt)
	if err != nil {
		return model.Cookies{}, helpers.NewInternalError(fmt.Errorf("can't create cookie: %w", err))
	}

	return cookie, nil
}
