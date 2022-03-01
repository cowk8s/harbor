package log

import (
	"context"
)

type loggerKey struct{}

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func GetLogger(ctx context.Context) *Logger {

	logger := ctx.Value(loggerKey{})

	return logger.(*Logger)
}
