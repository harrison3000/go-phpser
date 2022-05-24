package phpser

func (v *PhpValue) Any() any {
	switch v.pType {
	case TypeFloat:
		return v.num
	case TypeInt:
		return int(v.num)
	case TypeBool:
		return v.num != 0
	}

	return nil
}
