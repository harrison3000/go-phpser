package phpser

type phpType int8

const (
	typeNoExists phpType = iota
	typeNull
	typeBool

	typeArray
	typeObject

	typeFloat
	typeInt
	typeString

	//TODO references?
)

type mapKey struct {
	strKey string
	intKey int

	keyType phpType
}

type phpMapItem struct {
	key   mapKey
	value PhpValue
}

type PhpValue struct {
	str string
	arr []phpMapItem
	mmp map[mapKey]PhpValue
	num float64

	pType phpType
}
