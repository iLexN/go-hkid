package go_hkid

import "fmt"

type Hkid struct {
	part1 string
	part2 string
	part3 string
}

func NewHkidNil() *Hkid {
	return &Hkid{}
}

func NewHkid(part1 string, part2 string, part3 string) *Hkid {
	return &Hkid{part1: clearString(part1), part2: part2, part3: clearString(part3)}
}

func (hkid *Hkid) GetPart1() string {
	return hkid.part1
}

func (hkid *Hkid) GetPart2() string {
	return hkid.part2
}

func (hkid *Hkid) GetPart3() string {
	return hkid.part3
}

func (hkid *Hkid) Format() string {
	return fmt.Sprintf("%s%s(%s)", hkid.part1, hkid.part2, hkid.part3)
}
