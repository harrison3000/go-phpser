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
