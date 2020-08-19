package integers

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers:", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5}
	got := SumAll(numbers1, numbers2)
	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 9})
		want := []int{2, 9}
		checkSum(t, got, want)

	})
	t.Run("make the sums empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSum(t, got, want)
	})

}
