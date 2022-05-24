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
	assert.Equal(t, 5489, n.Value())

	n, e = Parse("b:1;")

	assert.Nil(t, e)
	assert.Equal(t, true, n.Value())

	n, e = Parse("b:0;")

	assert.Nil(t, e)
	assert.Equal(t, false, n.Value())

	n, e = Parse("d:2.565;")

	assert.Nil(t, e)
	assert.InDelta(t, 2.565, n.Value(), 0.00001)
}

func TestString(t *testing.T) {
	kv := map[string]string{
		`s:0:"";`:            "",
		`s:1:" ";`:           " ",
		"s:1:\"\x00\";":      "\x00",
		`s:10:"987456'""'";`: `987456'""'`,
	}

	for k, v := range kv {
		n, e := Parse(k)

		assert.Nil(t, e)
		assert.Equal(t, v, n.Value())
	}
}

func TestArray(t *testing.T) {
	n, e := Parse(`a:2:{s:3:"oxe";i:123;i:2;i:77;}`)
	_, _ = n, e
	//TODO actually implement this test
}

func TestObj(t *testing.T) {
	n, e := Parse("O:5:\"Objee\":2:{s:4:\"asas\";N;s:10:\"\x00Objee\x00exe\";N;}")
	_, _ = n, e
	//TODO actually implement this test
}
