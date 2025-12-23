package translator_to_postfix

import (
	"comp/parser"
	"strings"
)

type translator struct {
	*parser.Parser
	buf strings.Builder
}

func translate(input string) string {
	t := translator{
		Parser: parser.NewFromString(input),
	}
	t.expr()

	return t.buf.String()
}

func (t *translator) expr() {
	t.term()
	t.rest()
}

func (t *translator) rest() {
	switch t.Lookahead() {
	case "+":
		t.Match("+")
		t.term()
		t.buf.WriteString("+")
		t.rest()
	case "-":
		t.Match("-")
		t.term()
		t.buf.WriteString("-")
		t.rest()
	}
}

func (t *translator) term() {
	switch t.Lookahead() {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		tmp := t.Lookahead()
		t.Match(t.Lookahead())
		t.buf.WriteString(tmp)
	default:
		panic("invalid token")
	}
}
