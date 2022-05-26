package phpser

import (
	"fmt"
	"strconv"
)

func (v PhpValue) Value() any {
	switch {
	case v.IsFloat():
		return v.num
	case v.IsInt():
		return int(v.num)
	case v.IsBool():
		return v.num != 0
	case v.IsString():
		return v.str
	case v.pType == typeNoExists:
		return nil
	}

	return nil
}

func (v PhpValue) String() string {
	if v.str != "" {
		return v.str
	}

	switch v.pType {
	case typeNull:
		return "null"
	case typeBool:
		return strconv.FormatBool(v.num != 0)
	case typeInt, typeFloat:
		return fmt.Sprint(v.num)
	}

	return ""
}

func (v PhpValue) Get(key any) PhpValue {
	k := mkKey(key)
	return v.mmp[k]
}
