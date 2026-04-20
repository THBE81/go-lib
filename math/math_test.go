package math

import (
	"math"
	"testing"
)

func TestMapInt(t *testing.T) {
	var expected float64
	var result float64
	t.Run("5 from [0,10] to [0,100]", func(t *testing.T) {
		expected = 50.0
		result = MapRange(5.0, 0.0, 10.0, 0.0, 100.0)
		if result != expected {
			t.Errorf("MapInt() = %v, want %v", result, expected)
		}
		expected = 36.666
		result = MapRange(3, 2, 5, 15, 80)
		if math.Round(result) != math.Round(expected) {
			t.Errorf("MapInt() = %v, want %v", math.Round(result), math.Round(expected))
		}

	})

}
