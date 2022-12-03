package main

import (
	"context"
	"net/http"
	"quizON/internal/app/delivery"
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

	service := delivery.NewDelivery(db)

	r.Post("/login", service.Login)
	err := http.ListenAndServe(":"+config.GlobalConfig.Server.Port, r)
	if err != nil {
		logger.Fatalf("can't start server: %v", err)
	}
}
