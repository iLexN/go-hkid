package hkid

func CheckString(string string) *result {
	hkid, e := validatePatten(string)
	if e != nil {
		return NewHkidResultWithValid(newHkidNil(), false)
	}

	if getRemainder(hkid) == hkid.part3 {
		return NewHkidResultWithValid(hkid, true)
	}

	return NewHkidResultWithValid(hkid, false)
}

func CheckPart(part1 string, part2 string, part3 string) *result {
	hkid := newHkid(part1, part2, part3)
	return CheckString(hkid.Format())
}
