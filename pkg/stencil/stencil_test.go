package stencil

import (
	"testing"

	. "github.com/konzertmeister/konzertmeister/test"
)

func TestNewStencil(t *testing.T) {
	s, err := NewStencil("../../test/fixtures/basic_stencil.go.tmpl")
	if err != nil {
		t.Fatalf("Couldn't parse stencil: %v", err)
	}

	result := len(s.Arguments())
	expected := 1

	Eq(t, result, expected)
}
