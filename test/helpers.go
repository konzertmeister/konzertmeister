package test

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
)

var Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
	Level(zerolog.DebugLevel).
	With().
	Timestamp().
	Logger()

func Eq(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("Assertion failed, expected: %v, got: %v", expected, actual)
	}
}
