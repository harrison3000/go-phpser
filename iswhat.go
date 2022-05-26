package phpser

import "strings"

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

func (v PhpValue) IsArray() bool {
	return v.pType == TypeArray
}

func (v PhpValue) IsIterable() bool {
	return v.pType == TypeArray || v.pType == TypeObject
}

func (v PhpValue) Valid() bool {
	return v.pType != TypeInvalid
}

func (v PhpValue) IsObject() bool {
	return v.pType == TypeObject
}

func (v PhpValue) InstanceOf(class string, strict ...bool) bool {
	if v.pType != TypeObject {
		return false
	}

	if class == v.str {
		return true
	}

	if len(strict) > 0 && strict[0] {
		//if is strict and got here... nope
		return false
	}

	sep := strings.Split(v.str, `\`)

	cn := sep[len(sep)-1]

	return class == cn
}
