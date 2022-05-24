package phpser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {
	n, e := Parse("N;")

	assert.Nil(t, e)
	assert.True(t, n.IsNull())
}

func TestNumeric(t *testing.T) {
	n, e := Parse("i:5489;")

	assert.Nil(t, e)
	assert.Equal(t, 5489, n.Any())

	n, e = Parse("b:1;")

	assert.Nil(t, e)
	assert.Equal(t, true, n.Any())

	n, e = Parse("b:0;")

	assert.Nil(t, e)
	assert.Equal(t, false, n.Any())

	n, e = Parse("d:2.565;")

	assert.Nil(t, e)
	assert.InDelta(t, 2.565, n.Any(), 0.00001)
}
