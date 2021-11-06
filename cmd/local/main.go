package main

import (
	"github.com/paulhenri-l/golang-serverless/pkg/server"
	"net/http"
)

func main() {
	host := "127.0.0.1:3000"
	s := server.NewServer(host)
	// l.Info("Listening", zap.String("host", s.Addr))

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		//l.Error("Server error", zap.Error(errors.WithStack(err)))
		// close(quit)
	}
}
