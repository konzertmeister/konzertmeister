package stencil

import (
	"testing"
)

func TestNewStencil(t *testing.T) {
	s, err := NewStencil("../../test/fixtures/basic_stencil.go.tmpl")
	if err != nil {
		t.Fatalf("Couldn't parse stencil: %v", err)
	}

	result := len(s.Arguments())
	expected := 1

	if result != expected {
		t.Errorf("Invalid number of arguments parsed, expected: %d, got: %d",
			expected,
			result,
		)
	}
}
