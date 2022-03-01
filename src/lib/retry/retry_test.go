package retry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbort(t *testing.T) {
	assert := assert.New(t)

	e1 := Abort(nil)
	assert.Equal("retry abort", e1.Error())
}
