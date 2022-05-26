package phpser

func (v PhpValue) ForEach(iterator func(key, value PhpValue) bool) {
	if t := v.pType; t != typeArray && t != typeObject {
		iterator(PhpValue{}, v)
		return
	}

	for _, v := range v.arr {
		k := PhpValue{
			pType: v.key.keyType,
			str:   v.key.strKey,
			num:   float64(v.key.intKey),
		}

		keepGoing := iterator(k, v.value)
		if !keepGoing {
			break
		}
	}
}
