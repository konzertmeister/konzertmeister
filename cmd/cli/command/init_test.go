package command

import (
	"os"
	"testing"

	"github.com/konzertmeister/konzertmeister/test"
)

func TestMain(m *testing.M) {
	Logger = &test.Logger
	code := m.Run()
	os.Exit(code)
}
