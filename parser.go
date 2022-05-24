package phpser

import (
	"bufio"
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

	return
}
