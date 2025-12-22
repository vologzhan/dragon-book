package homework

import (
	"comp/2/lexer"
	"comp/2/parser"
	"testing"
)

func Test2(t *testing.T) {
	// S -> S ( S ) S | ε
	p := Parser2{
		parser.New(lexer.NewSymbols("(((())))")),
	}
	p.s()
}

type Parser2 struct {
	*parser.Parser
}

func (p *Parser2) s() {
	switch p.Lookahead() {
	case "(":
		p.Match("(")
		p.s()
		p.Match(")")
	}
}
