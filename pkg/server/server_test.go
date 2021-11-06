package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler()

	assert.NotNil(t, h)
}

func TestNewServer(t *testing.T) {
	s := NewServer("127.0.0.1:1234")

	assert.IsType(t, &http.Server{}, s)
	assert.Equal(t, "127.0.0.1:1234", s.Addr)
}
