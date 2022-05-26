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
		`s:16:"😀 😇 é 一";`:    "😀 😇 é 一",
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
	n := Parse("O:8:\"ns\\Objee\":2:{s:4:\"asas\";N;s:13:\"\x00ns\\Objee\x00exe\";N;}")

	assert.True(t, n.IsObject())
	assert.True(t, n.InstanceOf("Objee"))
	assert.False(t, n.InstanceOf("Objetu"))

	assert.True(t, n.InstanceOf(`ns\Objee`, true))
	assert.False(t, n.InstanceOf("Objee", true))

	//TODO implement more test
}

func TestBad(t *testing.T) {
	n := Parse("zzzz")

	v := n.Value()

	assert.False(t, n.Valid())
	assert.Error(t, v.(error))
}

func BenchmarkHugeVal(b *testing.B) {
	long := []byte(`a:2:{s:5:"teste";a:4:{i:0;s:19:"hellow mai friendis";i:1;s:24:"teste testancio da silva";s:12:"098203984098";i:6546546;i:2;i:123123123;}s:6:"teste2";a:5:{i:0;s:9:"AMD 4700S";i:1;s:19:"Zen, mas não Ryzen";s:6:"Distro";s:9:"ArchLinux";i:2;i:123123123;s:11:"outra array";a:15:{i:0;i:321;i:1;i:654;i:2;i:987;i:3;i:654;i:4;s:3:"asd";i:5;i:980;i:6;i:234235;i:7;i:234625;i:8;i:345;i:9;i:73;i:10;i:214;i:11;i:12;i:12;i:234;i:13;i:456;i:14;i:6234234234;}}}`)

	for i := 0; i < b.N; i++ {
		r := ParseBytes(long)
		_ = r
	}
}
