package go_hkid

type patterNotMatchError struct {
}

func NewPatterNotMatchError() *patterNotMatchError {
	return &patterNotMatchError{}
}

func (h patterNotMatchError) Error() string {
	return "Patter not match"
}
