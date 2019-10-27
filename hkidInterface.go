package go_hkid

type hkidInterface interface {
	GetPart1() string
	GetPart2() string
	GetPart3() string
	Format() string
}
