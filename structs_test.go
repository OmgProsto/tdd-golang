package golang_tdd

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeterRectangle(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaRectangle(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Area()
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestShapeAres(t *testing.T) {
	tests := []struct {
		Shape Shape
		Area  float64
	}{
		{Rectangle{10.0, 10.0}, 100},
		{Circle{10}, 314.1592653589793},
	}

	for _, shape := range tests {
		if shape.Shape.Area() != shape.Area {
			t.Errorf("got %.2f want %.2f", shape.Shape.Area(), shape.Area)
		}
	}
}
