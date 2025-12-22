package example

import (
	"comp/2/lexer"
	"comp/2/parser"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// Трансляция арифметических выражений в постфиксную запись
	// expr -> term rest
	//
	// rest -> +term {print("+")} rest
	//       | -term {print("-")} rest
	//       | ε
	//
	// term -> 0 {print("0)}
	//       | 1 {print("1)}
	//         ...
	//       | 9 {print("9)}

	input := "9-5+2"

	p := Translator{
		Parser: parser.New(lexer.NewSymbols(input)),
	}

	fmt.Printf("%s   ->   ", input)
	p.expr()
	fmt.Print("\n")
}

type Translator struct {
	*parser.Parser
}

func (t *Translator) expr() {
	t.term()
	t.rest()
}

func (t *Translator) rest() {
	switch t.Lookahead() {
	case "+":
		t.Match("+")
		t.term()
		fmt.Print("+")
		t.rest()
	case "-":
		t.Match("-")
		t.term()
		fmt.Print("-")
		t.rest()
	}
}

func (t *Translator) term() {
	switch t.Lookahead() {
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
		tmp := t.Lookahead()
		t.Match(t.Lookahead())
		fmt.Print(tmp)
	default:
		t.LogAndExit()
	}
}
