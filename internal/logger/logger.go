package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Logger is the implementation of Handler interface
type Handler struct {
	zerolog.Logger
}

// New instantiates bucky logger instance
func New(appname string, opts Options) (*Handler, error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger = logger.With().Str("app", appname).Logger()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	return &Handler{
		logger,
	}, nil
}
