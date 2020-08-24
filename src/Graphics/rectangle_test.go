package Graphics

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{20, 30}
	got := Perimeter(r)
	want := 100.0

	if got != want {
		t.Errorf("got: %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("长方形", func(t *testing.T) {
		rectangle := Rectangle{20, 30}
		want := 600.0

		checkArea(t, rectangle, want)
	})
	t.Run("圆形", func(t *testing.T) {
		circle := Circle{4}
		want := 16 * math.Pi

		checkArea(t, circle, want)
	})

}
