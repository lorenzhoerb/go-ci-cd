package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	a, b := 1, 2
	got := Add(a, b)
	want := 3

	if got != want {
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, got, want)
	}
}

func TestSub(t *testing.T) {
	a, b := 5, 2
	got := Sub(a, b)
	want := 3

	if got != want {
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, got, want)
	}
}
