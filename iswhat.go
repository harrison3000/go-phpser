package phpser

func (v PhpValue) IsNull() bool {
	return v.pType == TypeNull
}

func (v PhpValue) IsInt() bool {
	return v.pType == TypeInt
}

func (v PhpValue) IsBool() bool {
	return v.pType == TypeBool
}

func (v PhpValue) IsFloat() bool {
	return v.pType == TypeFloat
}

func (v PhpValue) IsString() bool {
	return v.pType == TypeString
}
