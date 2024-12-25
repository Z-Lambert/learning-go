package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)

		if result != 15 {
			t.Errorf("Expected 15, got %d", result)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)

		if result != 6 {
			t.Errorf("Expected 6, got %d", result)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	fmt.Println(result)
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
