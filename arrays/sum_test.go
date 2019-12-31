package main

import "testing"

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected int, collection []int) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %d but got %d given, %v", expected, got, collection)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		assertCorrectMessage(t, got, expected, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		assertCorrectMessage(t, got, expected, numbers)
	})
}
