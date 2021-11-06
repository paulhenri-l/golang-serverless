package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulhenri-l/golang-serverless/pkg/handlers"
	"github.com/treastech/logger"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewHandler(l *zap.Logger) http.Handler {
	r := chi.NewRouter()
	r.Use(logger.Logger(l))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

	r.Method("GET", "/", handlers.NewHelloHandler())

	return r
}

func NewServer(l *zap.Logger, host string) *http.Server {
	return &http.Server{
		Addr:         host,
		Handler:      NewHandler(l),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
}
