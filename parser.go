package phpser

import (
	"bufio"
	"bytes"
	"io"
	"reflect"
	"strconv"
	"strings"
)

func ParseBytes(serialized []byte) PhpValue {
	br := bytes.NewReader(serialized)
	return parse(br)
}

func Parse(serialized string) PhpValue {
	sr := strings.NewReader(serialized)
	return parse(sr)
}

func parse(r io.Reader) (ret PhpValue) {
	br := bufio.NewReader(r)
	defer func() {
		err := recover()
		if se, ok := err.(string); ok {
			ret.pType = typeNoExists
			ret.str = se //will allow the user to see what went wrong, if he so wants
		}
	}()

	ret = consume(br)
	return
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
		return PhpValue{pType: typeNull}
	}

	switch t {
	case 'b':
		ret.pType = typeBool
	case 'i':
		ret.pType = typeInt
	case 'd':
		ret.pType = typeFloat
	case 's':
		ret.pType = typeString
	case 'a':
		ret.pType = typeArray
	case 'O':
		ret.pType = typeObject
		//TODO support for serialized objects, "C"
	default:
		panic("unknown type")
	}

	expect(':')

	switch t {
	case 'b', 'i', 'd':
		ret.num = consumeFloat(r)
		return
	}

	len := consumeLen(r)

	if ret.pType == typeString || ret.pType == typeObject {
		expect('"')

		buf := make([]byte, len)
		qtdrd, err := io.ReadFull(r, buf)
		if qtdrd != len || err != nil {
			panic("truncated serialized value")
		}

		expect('"')

		ret.str = string(buf)

		if ret.pType == typeString {
			expect(';')
			return
		}

		expect(':')
		len = consumeLen(r)
	}

	//Only TypeArray and TypeObject should get here

	expect('{')

	ret.arr = make([]phpMapItem, 0, len)
	ret.mmp = make(map[mapKey]PhpValue, len)

	for i := 0; i < len; i++ {
		k := consume(r)
		v := consume(r)

		mk := mkKey(k.Value())
		if mk.keyType() == typeNoExists {
			panic("wrong type in array or object key")
		}

		ret.arr = append(ret.arr, phpMapItem{
			key:   mk,
			value: v,
		})
		ret.mmp[mk] = v
	}

	expect('}')
	return
}

func consumeLen(r *bufio.Reader) int {
	l, e := r.ReadString(':')
	if e != nil {
		panic("syntax error")
	}
	l = strings.TrimSuffix(l, ":")

	l64, e := strconv.ParseInt(l, 10, 0)
	if e != nil {
		panic("error getting length")
	}
	return int(l64)
}

func consumeFloat(r *bufio.Reader) float64 {
	v, e := r.ReadString(';')
	if e != nil {
		panic("syntax error")
	}

	v = strings.TrimSuffix(v, ";")

	num, e := strconv.ParseFloat(v, 64)
	if e != nil {
		panic("error converting numeric val")
	}

	return num
}

func mkKey(v any) (k mapKey) {
	vv := reflect.ValueOf(v)

	if vv.CanInt() {
		k.v = int(vv.Int())
	} else if vv.Type().Kind() == reflect.String {
		k.v = vv.String()
	}

	return k
}
