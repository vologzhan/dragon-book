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
		"empty": {
			"",
			[]interface{}(nil),
		},
		"expr": {
			"1+23",
			[]interface{}{
				newNum(1),
				newToken('+'),
				newNum(23),
			},
		},
		"expr_with_whitespaces": {
			" a -  bc ",
			[]interface{}{
				newWord(TagId, "a"),
				newToken('-'),
				newWord(TagId, "bc"),
			},
		},
		"comment_single_line": {
			"// comment",
			[]interface{}(nil),
		},
		"comment_single_line_after_expr": {
			"a+b // comment",
			[]interface{}{
				newWord(TagId, "a"),
				newToken('+'),
				newWord(TagId, "b"),
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			l := NewLexer(tt.input)
			assert.Equal(t, tt.expected, l.ScanAll())
		})
	}
}
