package phpser

import "strings"

func (v PhpValue) IsNull() bool {
	return v.pType == typeNull
}

func (v PhpValue) IsInt() bool {
	return v.pType == typeInt
}

func (v PhpValue) IsBool() bool {
	return v.pType == typeBool
}

func (v PhpValue) IsFloat() bool {
	return v.pType == typeFloat
}

func (v PhpValue) IsString() bool {
	return v.pType == typeString
}

func (v PhpValue) IsArray() bool {
	return v.pType == typeArray
}

func (v PhpValue) IsIterable() bool {
	return v.pType == typeArray || v.pType == typeObject
}

func (v PhpValue) Exists() bool {
	return v.pType != typeNoExists
}

func (v PhpValue) IsObject() bool {
	return v.pType == typeObject
}

func (v PhpValue) InstanceOf(class string, strict ...bool) bool {
	if v.pType != typeObject {
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
