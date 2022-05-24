package phpser

import (
	"bufio"
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
	}

	switch t {
	case 'b', 'i', 'd':
		expect(':')
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
	}

	return
}
