package hkid

import "fmt"

type hkid struct {
	part1 string
	part2 string
	part3 string
}

func newHkidNil() *hkid {
	return &hkid{}
}

func newHkid(part1 string, part2 string, part3 string) *hkid {
	return &hkid{part1: clearString(part1), part2: part2, part3: clearString(part3)}
}

func (hkid *hkid) GetPart1() string {
	return hkid.part1
}

func (hkid *hkid) GetPart2() string {
	return hkid.part2
}

func (hkid *hkid) GetPart3() string {
	return hkid.part3
}

func (hkid *hkid) Format() string {
	return fmt.Sprintf("%s%s(%s)", hkid.part1, hkid.part2, hkid.part3)
}
