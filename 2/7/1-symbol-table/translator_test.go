package symbol_table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	trans := NewTranslator("{ int x; char y; { bool y; x; y; } x; y; }")
	assert.Equal(t, "{ { x:int; y:bool; } x:int; y:char; }", trans.Translate())
}
