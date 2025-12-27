package union

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionTransitionDiagram(t *testing.T) {
	assert.Equal(t, []interface{}{
		KeyWord(IF),
		KeyWord(THEN),
		KeyWord(ELSE),
		Word("then2Else"),
		Number("123"),
		Float("3.14"),
		Float("1.23E+2"),
		Float("1E-2"),
		Float("1E3"),
		Relop(LT),
		Relop(LE),
		Relop(EQ),
		Relop(NE),
		Relop(GT),
		Relop(GE),
	}, parse("\tif then else then2Else 123 3.14 1.23E+2 1E-2 1E3< <= = <> > >=\n"))
}
