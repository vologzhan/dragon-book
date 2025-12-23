package symbol_table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolTable(t *testing.T) {
	assert.Equal(
		t,
		"{ { x:int; y:bool; } x:int; y:char; }",
		translate("{ int x; char y; { bool y; x; y; } x; y; }"),
	)
}
