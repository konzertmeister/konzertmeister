package command

import (
	"github.com/rs/zerolog"
)

var _logger *zerolog.Logger

func SetLogger(logger *zerolog.Logger) {
	_logger = logger
}

func setLogLevelDebug() {
	newLogger := _logger.Level(zerolog.DebugLevel)
	_logger = &newLogger
	_logger.Debug().Msg("Debug log activated")
}
