package main

import (
	"fmt"
	"github.com/paulhenri-l/golang-serverless/config"
	logger2 "github.com/paulhenri-l/golang-serverless/pkg/logger"
	"github.com/paulhenri-l/golang-serverless/pkg/server"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger

func init() {
	cfg := config.Get()
	l, err := logger2.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	logger = l
}

func main() {
	defer logger.Sync()
	defer handlePanic(logger)

	host := fmt.Sprintf("%s:%s", config.Get().AppHost, config.Get().AppPort)
	s := server.NewServer(logger, host)
	logger.Info("Listening", zap.String("host", s.Addr))

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Server error", zap.Error(errors.WithStack(err)))
	}
}

func handlePanic(l *zap.Logger) {
	err := recover()

	if err != nil {
		l.Error("panic", zap.String("error", fmt.Sprintf("%s", err)))
		l.Sync()
	}
}
