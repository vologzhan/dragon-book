package example

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
			"1+2",
			[]interface{}{
				Num{Token{TagNum}, 1},
				Token{'+'},
				Num{Token{TagNum}, 2},
			},
		},
		"expr_with_whitespaces": {
			"a -  b",
			[]interface{}{
				Word{Token{TagId}, "a"},
				Token{'-'},
				Word{Token{TagId}, "b"},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			l := NewLexer(tt.input)
			assert.Equal(t, tt.expected, l.ScanAll())
		})
	}
}
