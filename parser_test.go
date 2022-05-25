package phpser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNull(t *testing.T) {
	n := Parse("N;")

	assert.True(t, n.Valid())
	assert.True(t, n.IsNull())
}

func TestNumeric(t *testing.T) {
	n := Parse("i:5489;")

	assert.True(t, n.Valid())
	assert.Equal(t, 5489, n.Value())

	n = Parse("b:1;")

	assert.True(t, n.Valid())
	assert.Equal(t, true, n.Value())

	n = Parse("b:0;")

	assert.True(t, n.Valid())
	assert.Equal(t, false, n.Value())

	n = Parse("d:2.565;")

	assert.True(t, n.Valid())
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
		n := Parse(k)

		assert.True(t, n.Valid())
		assert.Equal(t, v, n.Value())
	}
}

func TestArray(t *testing.T) {
	n := Parse(`a:2:{s:3:"oxe";i:123;i:2;i:77;}`)
	_ = n
	//TODO actually implement this test
}

func TestObj(t *testing.T) {
	n := Parse("O:5:\"Objee\":2:{s:4:\"asas\";N;s:10:\"\x00Objee\x00exe\";N;}")
	_ = n
	//TODO actually implement this test
}

func TestBad(t *testing.T) {
	n := Parse("zzzz")

	v := n.Value()

	assert.False(t, n.Valid())
	assert.Error(t, v.(error))
}
