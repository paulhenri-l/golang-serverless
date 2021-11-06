package logger

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {
	l, err := NewLogger("info")

	assert.NoError(t, err)
	assert.IsType(t, &zap.Logger{}, l)
}
