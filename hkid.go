package hkid

// CheckString : Check HKID Format eg. CA182361(1).
func CheckString(string string) *result {
	hkid, e := validatePatten(string)
	//if errors.Is(e, newPatterNotMatchError()) {
	if e != nil {
		return newHkidResultWithValid(newHkidNil(), false)
	}

	if getRemainder(hkid) == hkid.part3 {
		return newHkidResultWithValid(hkid, true)
	}

	return newHkidResultWithValid(hkid, false)
}

// CheckPart : check HKID Format eg. CA182361(1).
// part1 = "CA"
// part2 = "182361"
// part3 = "1"
func CheckPart(part1 string, part2 string, part3 string) *result {
	hkid := newHkid(part1, part2, part3)
	return CheckString(hkid.Format())
}
