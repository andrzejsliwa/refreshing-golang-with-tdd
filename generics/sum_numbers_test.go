package sum_numbers

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		ints := map[string]int64{
			"first":  34,
			"second": 12,
		}

		gotInts := SumNumbers(ints)
		if gotInts != 46 {
			t.Errorf("Wanted %v, got %v", 46, gotInts)
		}
	})

	t.Run("floats", func(t *testing.T) {
		floats := map[string]float64{
			"first":  35.98,
			"second": 26.99,
		}

		gotFloats := SumNumbers(floats)
		if gotFloats != 62.97 {
			t.Errorf("Wanted %v, got %v", 62.97, gotFloats)
		}
	})
}
