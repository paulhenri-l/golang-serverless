package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello"))
	})

	return r
}

func NewServer(host string) *http.Server {
	return &http.Server{
		Addr:         host,
		Handler:      NewHandler(),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
}
