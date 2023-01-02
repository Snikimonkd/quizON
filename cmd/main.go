package main

import (
	"context"
	"net/http"
	"quizON/internal/app/delivery"
	cookieMiddleware "quizON/internal/app/middleware"
	"quizON/internal/config"
	"quizON/internal/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()
	db := config.ConnectToPostgres(ctx)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// мидалварь, которая проверяет куки
	checkCookieMiddleware := cookieMiddleware.NewCheckCookieMiddleware(db)
	service := delivery.NewDelivery(db)

	r.With(checkCookieMiddleware.CheckCookie).Post("/game", service.CreateGame)
	r.With(checkCookieMiddleware.CheckCookie).Get("/registrations", service.GetRegistrations)

	r.Post("/login", service.Login)
	r.Post("/registration", service.Register)

	r.Get("/games", service.GetGames)

	logger.Infof("server start at port: %v", config.GlobalConfig.Server.Port)
	err := http.ListenAndServe(":"+config.GlobalConfig.Server.Port, r)
	if err != nil {
		logger.Fatalf("can't start server: %v", err)
	}
}
