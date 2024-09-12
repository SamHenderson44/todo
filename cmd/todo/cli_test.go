package main

import (
	"strings"
	"testing"
)

func TestShowToDoOptions(t *testing.T) {
	t.Run("Calls ReadToDoText when 1 is entered", func(t *testing.T) {
		calledCount := 0
		ReadToDoText = mockFunction(&calledCount)

		input := strings.NewReader("1\n")

		ShowToDoOptions(input)
		assertError(t, 1, calledCount)
	})
	t.Run("Calls ShowToDos when 2 is entered", func(t *testing.T) {
		calledCount := 0
		ShowToDos = mockFunction(&calledCount)

		input := strings.NewReader("2\n")

		ShowToDoOptions(input)
		assertError(t, 1, calledCount)
	})
	t.Run("Checks for invalid input", func(t *testing.T) {
		calledCount := 0
		InvalidInput = mockFunction(&calledCount)

		strings.NewReader("someBadInput\n")

		InvalidInput()
		assertError(t, 1, calledCount)
	})
}

func assertError(t testing.TB, want int, got int) {
	t.Helper()

	if got != want {
		t.Errorf("expected calledCount to be %v, but got %v", got, want)
	}
}

var mockFunction = func(calledCount *int) func() {
	return func() {
		*calledCount++
	}
}
