package homework

import (
	"comp/2/lexer"
	"comp/2/parser"
	"testing"
)

func Test1(t *testing.T) {
	// S -> +SS | -SS | a

	p := Parser1{
		parser.New(lexer.NewSymbols("++++aaaaa")),
	}
	p.s()

	p = Parser1{
		parser.New(lexer.NewSymbols("+++--+-+-+-+-+++-a-aaaaaaaaaaaaaaaaaa")),
	}
	p.s()
}

type Parser1 struct {
	*parser.Parser
}

func (p *Parser1) s() {
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
		p.LogAndExit()
	}
}
