package logging

import (
	"context"

	"go.uber.org/zap"

	"github.com/kelseyhightower/envconfig"
)

var logger *zap.Logger

type loggerConfig struct {
	Format string `default:"MACHINE"`
}

const contextKey = "logger"

func init() {
	var env loggerConfig
	envconfig.MustProcess("LOG", &env)

	var log *zap.Logger

	if env.Format != "MACHINE" {
		log, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		logger = log
		logger.Info("Initialized development logger")
		return
	}

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = log
	logger.Info("Initialized production logger")

}

// GetLogger extract the logger from context if present
// or returns the default logger
func GetLogger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(contextKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}

// ContextWithLogger returns a context from the parent context with the logger attached
func ContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	if logger == nil {
		logger = GetLogger(ctx)
	}

	return context.WithValue(ctx, contextKey, logger)
}
