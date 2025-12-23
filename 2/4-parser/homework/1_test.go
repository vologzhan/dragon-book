package parser_homework

import (
	"comp/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserHomework1(t *testing.T) {
	// S -> +SS | -SS | a
	p := parser1{parser.New(newLexer("++++aaaaa"))}
	assert.NotPanics(t, p.s)

	p = parser1{parser.New(newLexer("+++--+-+-+-+-+++-a-aaaaaaaaaaaaaaaaaa"))}
	assert.NotPanics(t, p.s)
}

type parser1 struct {
	*parser.Parser
}

func (p *parser1) s() {
	switch p.Lookahead() {
	case "-":
		p.Match("-")
		p.s()
		p.s()
	case "+":
		p.Match("+")
		p.s()
		p.s()
	case "a":
		p.Match("a")
	default:
		panic("invalid token")
	}
}
