package phpser

type PhpType int8

const (
	TypeNull PhpType = iota
	TypeBool

	TypeArray

	TypeFloat
	TypeInt
	TypeString

	//TODO references... objects....
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
