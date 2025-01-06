package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of two arrays", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4, 5}
		arr2 := []int{3, 4, 5}

		got := SumAll(arr1, arr2)
		want := []int{15, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("run with normal arrays", func(t *testing.T) {
		got := sumAllTails([]int{1, 2}, []int{0, 9}, []int{1, 2, 3, 4, 5})
		want := []int{2, 9, 14}
		checkSums(t, got, want)
	})
	t.Run("run with an empty array", func(t *testing.T) {
		got := sumAllTails([]int{}, []int{0, 9}, []int{1, 2, 3, 4, 5})
		want := []int{0, 9, 14}

		checkSums(t, got, want)
	})
}
