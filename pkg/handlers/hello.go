package handlers

import (
	"net/http"
)

type HelloHandler struct {}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

func (p *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	_, _ = rw.Write([]byte("Hello"))
}
