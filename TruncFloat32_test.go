package compare

import (
	"testing"
)

func TestTruncFloat32(t *testing.T) {
	t.Run("Test cases with positive values", func(t *testing.T) {
		value := float32(3.14159)
		position := 2
		expected := float32(3.14)
		result := TruncFloat32(value, position)
		if result != expected {
			t.Fatalf("TruncFloat64(%f, %d) = %f, expected %f", value, position, result, expected)
		}
	})

	t.Run("Test cases with negative values", func(t *testing.T) {
		// Test cases with negative values
		value := float32(-5.678)
		position := 1
		expected := float32(-5.6)
		result := TruncFloat32(value, position)
		if result != expected {
			t.Fatalf("TruncFloat32(%f, %d) = %f, expected %f", value, position, result, expected)
		}

	})

	t.Run("Test cases with zero value", func(t *testing.T) {
		value := float32(0)
		position := 3
		expected := float32(0)
		result := TruncFloat32(value, position)
		if result != expected {
			t.Fatalf("TruncFloat32(%f, %d) = %f, expected %f", value, position, result, expected)
		}
	})
}
