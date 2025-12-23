package parser_homework

import (
	"comp/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserHomework2(t *testing.T) {
	// S -> S ( S ) S | ε
	p := parser2{parser.New(newLexer("(((())))"))}
	assert.NotPanics(t, p.s)
}

type parser2 struct {
	*parser.Parser
}

func (p *parser2) s() {
	switch p.Lookahead() {
	case "(":
		p.Match("(")
		p.s()
		p.Match(")")
	case ")":
		return
	default:
		panic("invalid token")
	}
}
