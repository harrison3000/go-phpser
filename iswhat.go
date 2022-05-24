package phpser

func (v *PhpValue) IsNull() bool {
	return v.pType == TypeNull
}
