package phpser

type PhpType int8

const (
	TypeNull PhpType = iota
	TypeBool

	TypeArray
	TypeObject

	TypeFloat
	TypeInt
	TypeString

	//TODO references?
)

type mapKey struct {
	keyType PhpType

	intKey int
	strKey string
}

type PhpMapItem struct {
	key mapKey

	Value PhpValue
}

type PhpValue struct {
	pType PhpType

	str string
	num float64
	arr []PhpMapItem
}
