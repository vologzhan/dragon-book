package parser_example

import (
	"comp/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserExample(t *testing.T) {
	// Грамматика для некоторых выражений C и Java
	// stmt -> expr;
	//      |  if ( expr ) stmt
	//      |  for ( optexpr ; optexpr ; optexpr ) stmt
	//      |  other
	//
	// optexpr -> ε
	//         |  expr
	p := &Parser{parser.NewFromString("for ( ; expr ; expr ) other")}
	assert.NotPanics(t, p.stmt)
}

type Parser struct {
	*parser.Parser
}

func (p *Parser) stmt() {
	switch p.Lookahead() {
	case "expr":
		p.Match("expr")
		p.Match(";")
	case "if":
		p.Match("if")
		p.Match("(")
		p.Match("expr")
		p.Match(")")
		p.stmt()
	case "for":
		p.Match("for")
		p.Match("(")
		p.MatchOptional()
		p.Match(";")
		p.MatchOptional()
		p.Match(";")
		p.MatchOptional()
		p.Match(")")
		p.stmt()
	case "other":
		p.Match("other")
	default:
		panic("invalid token")
	}
}
