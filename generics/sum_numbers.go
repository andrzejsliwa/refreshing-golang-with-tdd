package sum_numbers

import "fmt"

type Number interface {
	int64 | float64
}

type Age interface {
	int | float64
}

func newGenericFunc[A Age](age A) string {
	return fmt.Sprint(age)
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
