package main
import (
	"testing"
)

func TestCalculator(t *testing.T) {
	cal := Calculator{}

	t.Run("Add", func(t *testing.T) {
		result := cal.Add(5, 3)
		expected := 8
		if result != expected {
			t.Errorf("Add(5, 3) = %d; want %d", result, expected)
		}
	})

	t.Run("Subtract", func(t *testing.T) {
		result := cal.Subtract(5, 3)
		expected := 2
		if result != expected {
			t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
		}
	})

	t.Run("Multiply", func(t *testing.T) {
		result := cal.Multiply(5, 3)
		expected := 15
		if result != expected {
			t.Errorf("Multiply(5, 3) = %d; want %d", result, expected)
		}
	})

	t.Run("Divide", func(t *testing.T) {
		result, err := cal.Divide(6, 3)
		expected := 2
		if err != nil {
			t.Errorf("Divide(6, 3) returned error: %v", err)
		}
		if result != expected {
			t.Errorf("Divide(6, 3) = %d; want %d", result, expected)
		}
	})

	t.Run("Divide by zero", func(t *testing.T) {
		_, err := cal.Divide(5, 0)
		if err == nil {
			t.Error("Divide(5, 0) did not return an error")
		}
	})
}	