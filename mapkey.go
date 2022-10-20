package phpser

func (k mapKey) String() string {
	s, _ := k.v.(string)

	return s
}

func (k mapKey) Int() int {
	i, _ := k.v.(int)

	return i
}

func (k mapKey) keyType() phpType {
	switch k.v.(type) {
	case string:
		return typeString
	case int:
		return typeInt
	default:
		return typeNoExists
	}
}

func (k mapKey) toVal() PhpValue {
	v := PhpValue{
		pType: k.keyType(),
		str:   k.String(),
		num:   float64(k.Int()),
	}

	return v
}
