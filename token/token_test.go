package token

import (
	"testing"
)

// Test of add function
func TestAdd(t *testing.T) {
	if (ParseIllegal('=') != Token{Type: ASSIGN, Literal: "="}) {
		t.Error("NOOO")
	}
}
