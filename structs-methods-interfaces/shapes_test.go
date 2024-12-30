package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 5.0}
	got := rectangle.Perimeter()
	want := 30.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func ExamplePerimeter() {
// 	rectangle := Rectangle{10.0, 5.0}
// 	perimeter := rectangle.Perimeter()
// 	fmt.Printf("%.2f", perimeter)
// 	// Output: 30.00
// }

// func BenchmarkPerimeter(b *testing.B) {
// 	rectangle := Rectangle{10.0, 5.0}
// 	for i := 0; i < b.N; i++ {
// 		rectangle.Perimeter()
// 	}
// }

// AREA

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 5.0}
		got := rectangle.Area()
		want := 50.00

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}

// func ExampleArea() {
// 	rectangle := Rectangle{10.0, 5.0}
// 	area := Area(rectangle)
// 	fmt.Printf("%.2f", area)
// 	// Output: 50.00
// }
//
// func BenchmarkArea(b *testing.B) {
// 	rectangle := Rectangle{10.0, 5.0}
// 	for i := 0; i < b.N; i++ {
// 		Area(rectangle)
// 	}
// }
