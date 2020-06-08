package test

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
	Level(zerolog.DebugLevel).
	With().
	Timestamp().
	Logger()
