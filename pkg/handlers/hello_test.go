package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewHelloHandler(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	rw := httptest.NewRecorder()

	handler := NewHelloHandler()
	handler.ServeHTTP(rw, r)

	assert.Equal(t, "Hello", rw.Body.String())
	assert.Equal(t, 200, rw.Code)
}

