package parser

import "comp/lexer"

type Parser struct {
	lexer     lexer.Interface
	lookahead string
}

func New(l lexer.Interface) *Parser {
	p := &Parser{
		lexer: l,
	}
	p.next()

	return p
}

func NewFromString(buf string) *Parser {
	return New(lexer.New(buf))
}

func (p *Parser) Lookahead() string {
	return p.lookahead
}

func (p *Parser) Match(terminal string) {
	if p.Lookahead() != terminal {
		panic("invalid token")
	}

	p.next()
}

func (p *Parser) MatchAny() {
	if p.lookahead == "" {
		panic("invalid token")
	}
	p.next()
}

func (p *Parser) MatchOptional() {
	if p.Lookahead() == "expr" {
		p.Match("expr")
	}
}

func (p *Parser) next() {
	p.lookahead = p.lexer.Next()
}
