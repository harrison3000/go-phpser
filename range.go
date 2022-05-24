package phpser

func (v PhpValue) ForEach(iterator func(key, value PhpValue) bool) {
	if t := v.pType; t != TypeArray && t != TypeObject {
		iterator(PhpValue{}, v)
		return
	}

	for _, v := range v.arr {
		k := PhpValue{
			pType: v.key.keyType,
			str:   v.key.strKey,
			num:   float64(v.key.intKey),
		}

		keepGoing := iterator(k, v.Value)
		if !keepGoing {
			break
		}
	}
}
