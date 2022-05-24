package phpser

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Parse(serialized string) (PhpValue, error) {
	sr := strings.NewReader(serialized)
	br := bufio.NewReader(sr)

	//TODO recover from panic and actually return error

	return consume(br), nil
}

func consume(r *bufio.Reader) (ret PhpValue) {
	expect := func(b byte) {
		bb, e := r.ReadByte()
		if e != nil || bb != b {
			panic("unexpected byte")
		}
	}

	t, e := r.ReadByte()
	if e != nil {
		panic("syntax error")
	}

	if t == 'N' {
		expect(';')
		return PhpValue{}
	}

	switch t {
	case 'b':
		ret.pType = TypeBool
	case 'i':
		ret.pType = TypeInt
	case 'd':
		ret.pType = TypeFloat
	case 's':
		ret.pType = TypeString
	case 'a':
		ret.pType = TypeArray
	case 'O':
		ret.pType = TypeObject
	default:
		panic("unknown type")
	}

	expect(':')

	var len int

	switch t {
	case 'b', 'i', 'd':
		v, e := r.ReadString(';')
		if e != nil {
			panic("syntax error")
		}

		v = strings.TrimSuffix(v, ";")

		ret.num, e = strconv.ParseFloat(v, 64)
		if e != nil {
			panic("error converting numeric val")
		}
		return
	case 's', 'a', 'O':
		l, e := r.ReadString(':')
		if e != nil {
			panic("syntax error")
		}
		l = strings.TrimSuffix(l, ":")

		l64, e := strconv.ParseInt(l, 10, 0)
		if e != nil {
			panic("error getting length")
		}
		len = int(l64)
	}

	//TODO object parsing

	if ret.pType == TypeString {
		expect('"')

		buf := make([]byte, len)
		qtdrd, err := io.ReadFull(r, buf)
		if qtdrd != len || err != nil {
			panic("truncated serialized value")
		}

		expect('"')
		expect(';')

		ret.str = string(buf)
		return
	}

	if ret.pType == TypeArray {
		expect('{')

		for i := 0; i < len; i++ {
			k := consume(r)
			v := consume(r)

			ret.arr = append(ret.arr, PhpMapItem{
				key:   k.mkKey(),
				Value: v,
			})
		}

		expect('}')
		return
	}

	return
}

func (v PhpValue) mkKey() mapKey {
	var k mapKey

	switch v.pType {
	case TypeInt:
		k.keyType = TypeInt
		k.intKey = int(v.num)
	case TypeString:
		k.keyType = TypeString
		k.strKey = v.str
	default:
		panic("wrong map key")
	}

	return k
}
