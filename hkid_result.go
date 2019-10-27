package go_hkid

type HkidResult struct {
	hkid  *Hkid
	Valid bool
}

func (h *HkidResult) GetPart1() string {
	return h.hkid.GetPart1()
}

func (h *HkidResult) GetPart2() string {
	return h.hkid.GetPart2()
}

func (h *HkidResult) GetPart3() string {
	return h.hkid.GetPart3()
}

func (h *HkidResult) Format() string {
	return h.hkid.Format()
}

func NewHkidResultWithValid(hkid *Hkid, valid bool) *HkidResult {
	return &HkidResult{hkid: hkid, Valid: valid}
}
