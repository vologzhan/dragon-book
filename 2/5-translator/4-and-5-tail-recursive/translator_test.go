package tail_recursive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTailRecursive(t *testing.T) {
	// Устранение оконечной рекурсии (tail recursive) из 2.5.3
	// Если последняя инструкция процедуры - вызов той же самой процедуры, то это называется оконечной рекурсией.
	assert.Equal(t, "95-2+", translate("9-5+2"))
}
