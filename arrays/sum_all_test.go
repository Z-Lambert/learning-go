package main

import (
	"fmt"
	"testing"
)

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 2})
	want := 9

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2, 3}, []int{1, 2})
	fmt.Println(sum)
	// Output: 9
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{1, 2})
	}
}
