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

type PhpMapItem struct {
	keyType PhpType

	intKey int
	strKey string

	Value PhpValue
}

type PhpValue struct {
	pType PhpType

	str string
	num float64
	arr []PhpMapItem
}
