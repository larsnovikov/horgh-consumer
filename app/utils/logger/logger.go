package logger

import (
	"context"
	"go.uber.org/zap"
)

const loggerKey = "logger"

func Set(ctx context.Context) (context.Context, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return ctx, err
	}
	defer logger.Sync()

	return context.WithValue(ctx, loggerKey, logger), nil
}

func Get(ctx context.Context) *zap.Logger {
	return ctx.Value(loggerKey).(*zap.Logger)
}
