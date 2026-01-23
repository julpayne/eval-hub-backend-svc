package logging

import (
	"log/slog"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
	"go.uber.org/zap/zapcore"
)

type ShutdownFunc func() error

// NewLogger creates and returns a new structured logger using zap as the underlying
// logging implementation, wrapped with slog's interface. The logger is configured
// with production settings and ISO8601 time encoding for consistent log formatting.
//
// Returns:
//   - *slog.Logger: A structured logger instance that can be used throughout the application
//   - error: An error if the logger could not be initialized
func NewLogger() (*slog.Logger, ShutdownFunc, error) {
	var logConfig zap.Config
	logConfig = zap.NewProductionConfig()
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLog, err := logConfig.Build()
	if err != nil {
		return nil, nil, err
	}
	f := newShutdownFunc(zapLog.Core())
	return slog.New(zapslog.NewHandler(zapLog.Core())), f, nil
}

func FallbackLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func newShutdownFunc(core zapcore.Core) ShutdownFunc {
	return func() error {
		return core.Sync()
	}
}
