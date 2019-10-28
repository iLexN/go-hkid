package hkid

type patterNotMatchError struct {
}

func newPatterNotMatchError() *patterNotMatchError {
	return &patterNotMatchError{}
}

func (h patterNotMatchError) Error() string {
	return "Patter not match"
}
