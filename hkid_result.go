package hkid

type result struct {
	hkid  *hkid
	Valid bool
}

func (h *result) GetPart1() string {
	return h.hkid.GetPart1()
}

func (h *result) GetPart2() string {
	return h.hkid.GetPart2()
}

func (h *result) GetPart3() string {
	return h.hkid.GetPart3()
}

func (h *result) Format() string {
	return h.hkid.Format()
}

func newHkidResultWithValid(hkid *hkid, valid bool) *result {
	return &result{hkid: hkid, Valid: valid}
}
