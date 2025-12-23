package parser_homework

import (
	"comp/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserHomework3(t *testing.T) {
	// S -> 0 S 1 | 0 1
	p := parser3{
		parser.New(newLexer("01")),
	}
	assert.NotPanics(t, p.s)
}

type parser3 struct {
	*parser.Parser
}

func (p *parser3) s() {
	switch p.Lookahead() {
	case "0":
		p.Match("0")
		p.s()
		p.Match("1")
	case "1":
		return
	default:
		panic("invalid token")
	}
}
