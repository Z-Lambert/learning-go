package main

import (
	"fmt"
	"reflect"
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

func ExampleRectangle_Perimeter() {
	rectangle := Rectangle{10.0, 5.0}
	perimeter := rectangle.Perimeter()
	fmt.Printf("%.2f", perimeter)
	// Output: 30.00
}

func BenchmarkPerimeter(b *testing.B) {
	rectangle := Rectangle{10.0, 5.0}
	for i := 0; i < b.N; i++ {
		rectangle.Perimeter()
	}
}

// AREA

var areaTests = []struct {
	shape Shape
	want  float64
}{
	{shape: Rectangle{Width: 10.0, Height: 5.0}, want: 50.0},
	{shape: Circle{Radius: 10.0}, want: 314.1592653589793},
	{shape: Triangle{Base: 5.0, Height: 6.0}, want: 15.0},
}

func TestArea(t *testing.T) {
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.want)
		}
	}
}

func exampleArea(shape Shape) {
	fmt.Printf("%.2f", shape.Area())
}

func ExampleCircle_Area() {
	circle := Circle{10.0}
	exampleArea(circle)
	// Output: 314.16
}

func ExampleRectangle_Area() {
	rectangle := Rectangle{10.0, 5.0}
	exampleArea(rectangle)
	// Output: 50.00
}

func ExampleTriangle_Area() {
	triangle := Triangle{5.0, 6.0}
	exampleArea(triangle)
	// Output: 15.00
}

func BenchmarkShapes(b *testing.B) {
	for _, tt := range areaTests {
		b.Run(reflect.TypeOf(tt.shape).String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tt.shape.Area()
			}
		})
	}
}
