package phpser

import (
	"fmt"
	"strconv"
	"time"
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
	case v.InstanceOf("DateTime"), v.InstanceOf("DateTimeImmutable"):
		return v.Time()
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

func (v PhpValue) Int() int {
	//TODO improve this!
	return int(v.num)
}

func (v PhpValue) Get(key any) PhpValue {
	k := mkKey(key)
	return v.mmp[k]
}

func (v PhpValue) Time() time.Time {
	if !v.InstanceOf("DateTime") && !v.InstanceOf("DateTimeImmutable") {
		return time.Time{}
	}

	const format = "2006-01-02 15:04:05.99999999Z07:00"

	date := v.Get("date").String()
	tztype := v.Get("timezone_type").Int()
	tz := v.Get("timezone").String()

	switch tztype {
	case 1:
		date += tz
	}

	t, _ := time.Parse(format, date)

	return t
}
