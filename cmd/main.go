package main

import (
	"net/http"
	"quizON/interanl/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	configPath = ""
)

func main() {
	con, err := config.NewConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
