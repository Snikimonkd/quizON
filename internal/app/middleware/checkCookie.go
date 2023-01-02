package middleware

import (
	"context"
	"fmt"
	"net/http"
	"quizON/internal/app/helpers"
	"quizON/internal/app/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type key int

const CtxUserID key = -1

type CheckCookieRepository interface {
	CheckCookie(ctx context.Context, value uuid.UUID) (int32, error)
}

type checkCookieMiddleware struct {
	checkCookieRepository CheckCookieRepository
}

func NewCheckCookieMiddleware(db *pgx.Conn) *checkCookieMiddleware {
	checkCookieRepository := repository.NewCheckCookieRepository(db)

	return &checkCookieMiddleware{
		checkCookieRepository: checkCookieRepository,
	}
}

func (c *checkCookieMiddleware) CheckCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cookie, err := r.Cookie("token")
		if err != nil {
			resp := helpers.NewHttpError(http.StatusUnauthorized, fmt.Errorf("can't find cookie in request: %v", err), helpers.UnauthorizedError)
			helpers.HandleError(w, resp)
			return
		}

		value, err := uuid.Parse(cookie.Value)
		if err != nil {
			resp := helpers.NewHttpError(http.StatusUnauthorized, fmt.Errorf("cna't parse cookie uuid: %v", err), helpers.UnauthorizedError)
			helpers.HandleError(w, resp)
			return
		}

		id, err := c.checkCookieRepository.CheckCookie(ctx, value)
		if errors.Is(err, repository.NotFoundError) {
			resp := helpers.NewHttpError(http.StatusUnauthorized, nil, helpers.UnauthorizedError)
			helpers.HandleError(w, resp)
			return
		}
		if err != nil {
			resp := helpers.NewHttpError(http.StatusInternalServerError, fmt.Errorf("cna't check cookie in db: %v", err), helpers.UnauthorizedError)
			helpers.HandleError(w, resp)
			return
		}

		ctx = context.WithValue(ctx,
			CtxUserID,
			id,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
