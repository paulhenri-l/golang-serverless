package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/handlerfunc"
	"github.com/paulhenri-l/golang-serverless/pkg/server"
	"net/http"
)

var httpHandler *handlerfunc.HandlerFuncAdapterV2

func init() {
	router := server.NewHandler()

	httpHandler = handlerfunc.NewV2(func(writer http.ResponseWriter, request *http.Request) {
		router.ServeHTTP(writer, request)
	})
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return httpHandler.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
