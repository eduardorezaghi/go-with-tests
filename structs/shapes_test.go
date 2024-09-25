package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Shape interface {
	Area() float64
}


func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	assert.Equal(t, want, got, "got %.2f want %.2f", got, want)
}


func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assert.Equal(t, want, got, "got %.2f want %.2f", got, want)
	}

	t.Run("Test Area with positive values", func(t *testing.T) {
		rectangle := Rectangle{2.0, 2.0}
		checkArea(t, rectangle, 4.0)
	});
	t.Run("Test Area with negative values", func(t *testing.T) {
		rectangle := Rectangle{-1.0, -1.0}
		checkArea(t, rectangle, 1.0)
	});

	t.Run("Test Area with circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	});
}