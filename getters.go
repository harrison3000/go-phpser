package phpser

import "errors"

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
	case v.pType == TypeInvalid:
		return errors.New(v.str)
	}

	return nil
}
