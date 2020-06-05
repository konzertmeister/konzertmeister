package main

import (
	"os"

	"github.com/rs/zerolog"
)

func initLog() *zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Logger()

	return &logger
}
