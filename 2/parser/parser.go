package parser

import (
	"comp/2/lexer"
	"fmt"
	"os"
)

func New(lexer lexer.Lexer) *Parser {
	p := &Parser{
		Lexer: lexer,
	}
	p.next()

	return p
}

type Parser struct {
	Lexer     lexer.Lexer
	lookahead string
}

func (p *Parser) Lookahead() string {
	return p.lookahead
}

func (p *Parser) Match(terminal string) {
	if p.Lookahead() != terminal {
		p.LogAndExit()
	}

	p.next()
}

func (p *Parser) MatchOptional() {
	if p.Lookahead() == "expr" {
		p.Match("expr")
	}
}

func (p *Parser) LogAndExit() {
	// todo табы при выводе ^
	fmt.Printf(`invalid syntax, pos: %d
%s
%*s
`, p.Lexer.Pos(), p.Lexer.Line(), p.Lexer.Pos()+1, "^")

	os.Exit(1)
}

func (p *Parser) next() {
	p.lookahead = p.Lexer.Next()
}
