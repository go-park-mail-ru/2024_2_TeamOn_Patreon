package validate

// Валидация для постов
const (
	MinLenTitle int = 3
	MaxLenTitle int = 100

	MaxLenContent int = 1000

	MinLayer        = 0
	MinSubLayer int = 1
	MaxLayer        = 3

	MinLenSubTitle int = 3
	MaxLenSubTitle int = 50

	MinSubCost int = 0
	MaxSubCost int = 1_000_000

	MinLenSubDescription int = 0
	MaxLenSubDescription int = 1000
)
