package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	for name, tt := range map[string]struct {
		input    string
		expected []interface{}
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
			"1 // comment",
			[]interface{}{
				newNum(1),
			},
		},
		"comment_multiline": {
			`4/*
*/ / /*
*/2`,
			[]interface{}{
				newNum(4),
				newToken('/'),
				newNum(2),
			},
		},
		"float": {
			"3.14",
			[]interface{}{
				newFloat(3.14),
			},
		},
		"float_suffix": {
			"2.",
			[]interface{}{
				newFloat(2.),
			},
		},
		"float_prefix": {
			".5",
			[]interface{}{
				newFloat(.5),
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			l := NewLexer(tt.input)
			assert.Equal(t, tt.expected, l.ScanAll())
		})
	}
}
