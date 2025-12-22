package example

import (
	"comp/2/lexer"
	"comp/2/parser"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// Устранение оконечной рекурсии (tail recursive) из 2.5.3
	// Если последняя инструкция процедуры - вызов той же самой процедуры, то это называется оконечной рекурсией.

	input := "9-5+2"

	p := Parser{
		Parser: parser.New(lexer.NewSymbols(input)),
	}

	fmt.Printf("%s   ->   ", input)
	p.expr()
	fmt.Print("\n")
}

type Parser struct {
	*parser.Parser
}

func (p *Parser) expr() {
	p.term()

	for {
		switch p.Lookahead() {
		case "+":
			p.Match("+")
			p.term()
			fmt.Print("+")
		case "-":
			p.Match("-")
			p.term()
			fmt.Print("-")
		default:
			return
		}
	}
}

func (p *Parser) term() {
	switch p.Lookahead() {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		t := p.Lookahead()
		p.Match(p.Lookahead())
		fmt.Print(t)
	default:
		p.LogAndExit()
	}
}
