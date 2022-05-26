package phpser

type PhpType int8

const (
	TypeInvalid PhpType = iota
	TypeNull
	TypeBool

	TypeArray
	TypeObject

	TypeFloat
	TypeInt
	TypeString

	//TODO references?
)

type mapKey struct {
	strKey string
	intKey int

	keyType PhpType
}

type PhpMapItem struct {
	key mapKey

	Value PhpValue
}

type PhpValue struct {
	str string
	arr []PhpMapItem
	mmp map[mapKey]PhpValue
	num float64

	pType PhpType
}
