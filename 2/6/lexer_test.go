package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	for name, tt := range map[string]struct {
		input    string
		expected interface{}
	}{
		"expr": {
			"1+23",
			[]interface{}{
				Num{Token{TagNum}, 1},
				Token{'+'},
				Num{Token{TagNum}, 23},
			},
		},
		"expr_with_whitespaces": {
			" a -  bc ",
			[]interface{}{
				Word{Token{TagId}, "a"},
				Token{'-'},
				Word{Token{TagId}, "bc"},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			l := NewLexer(tt.input)
			assert.Equal(t, tt.expected, l.ScanAll())
		})
	}
}
