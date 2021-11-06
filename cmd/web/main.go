package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/paulhenri-l/golang-serverless/config"
	logger2 "github.com/paulhenri-l/golang-serverless/pkg/logger"
	"github.com/paulhenri-l/golang-serverless/pkg/server"
	"go.uber.org/zap"
	"net/http"
)

var logger *zap.Logger
var httpHandler *handlerfunc.HandlerFuncAdapterV2

func init() {
	cfg := config.Get()
	l, err := logger2.NewLogger(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	logger  = l
	router := server.NewHandler(l)

	httpHandler = handlerfunc.NewV2(func(writer http.ResponseWriter, request *http.Request) {
		router.ServeHTTP(writer, request)
	})
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return httpHandler.ProxyWithContext(ctx, req)
}

func main() {
	defer logger.Sync()
	defer handlePanic(logger)

	lambda.Start(Handler)
}

func handlePanic(l *zap.Logger) {
	err := recover()

	if err != nil {
		l.Error("panic", zap.String("error", fmt.Sprintf("%s", err)))
		l.Sync()
	}
}
