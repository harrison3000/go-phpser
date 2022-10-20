package phpser

func (v PhpValue) ForEach(iterator func(key, value PhpValue) bool) {
	if t := v.pType; t != typeArray && t != typeObject {
		iterator(PhpValue{}, v)
		return
	}

	for _, v := range v.arr {
		k := v.key.toVal()

		keepGoing := iterator(k, v.value)
		if !keepGoing {
			break
		}
	}
}
