package translator_to_postfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslatorToPostfix(t *testing.T) {
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
	assert.Equal(t, "95-2+", translate("9-5+2"))
}
