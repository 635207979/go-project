package Graphics

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(20.0, 30.0)
	want := 100.0

	if got != want {
		t.Errorf("got: %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(20, 30)
	want := 600.0

	if got != want {
		t.Errorf("got: %.2f want %.2f", got, want)
	}
}
