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
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rect, want)
	})
	t.Run("circles", func(t *testing.T) {
		cir := Circle{10}
		want := math.Pi * 100
		checkArea(t, cir, want)

	})
}
