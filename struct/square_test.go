package _struct

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0
	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got '%.2f' want '%.2f'", shape, got, want)
		}
	}
	areaTest := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12,6},72},
		{Circle{10},math.Pi*100},
		{Triangle{12,6},36},
	}
	for _, tt := range areaTest {
		checkArea(t, tt.shape, tt.want)
	}
}
