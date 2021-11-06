package logger

import (
	"github.com/paulhenri-l/goenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) (*zap.Logger, error) {
	var cfg zap.Config
	l := zapcore.InfoLevel

	if err := l.UnmarshalText([]byte(level)); err != nil {
		return nil, err
	}

	if goenv.AppEnv() == goenv.Local || goenv.AppEnv() == "testing"{
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
		cfg.DisableStacktrace = false
	}

	cfg.Level = zap.NewAtomicLevelAt(l)

	return cfg.Build()
}
