package homework

import (
	"comp/2/lexer"
	"comp/2/parser"
	"testing"
)

func Test3(t *testing.T) {
	// S -> 0 S 1 | 0 1
	p := Parser3{
		parser.New(lexer.NewSymbols("01")),
	}
	p.s()
}

type Parser3 struct {
	*parser.Parser
}

func (p *Parser3) s() {
	switch p.Lookahead() {
	case "0":
		p.Match("0")
		p.s()
		p.Match("1")
	}
}
