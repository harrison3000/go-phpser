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
	v any
}

type phpMapItem struct {
	key   mapKey
	value PhpValue
}

type PhpValue struct {
	str string
	mmp map[mapKey]PhpValue
	arr []phpMapItem
	num float64

	pType phpType
}
