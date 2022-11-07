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

type mapItem struct {
	key   mapKey
	value PhpValue
}

type PhpValue struct {
	str string
	mmp map[mapKey]PhpValue
	arr []mapItem //Used for iterating in the right order, arrays in php are ordered maps
	num float64

	pType phpType
}
