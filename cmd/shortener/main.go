package main

import (
	"net/http"

	"github.com/Arseniy-Polyakov/GoCourse/cmd/config"
	handlers "github.com/Arseniy-Polyakov/GoCourse/internal/controller/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfgMain := config.NewConfig()
	parseFlags(cfgMain)
	// http.HandleFunc("/", handlers.HandlerPost)
	// http.HandleFunc("/{shortLink}", handlers.HandlerGet)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", handlers.HandlerPost)
	r.Get("/{shortLink}", handlers.HandlerGet)
	http.ListenAndServe(":8080", r)
}
