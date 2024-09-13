package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

func TestSelectNewToDoStatus(t *testing.T) {
	t.Run("returns correct input with spaces trimmed", func(t *testing.T) {
		input := strings.NewReader("1 \n")
		want := "1"

		got := SelectNewToDoStatus(input)

		if got != want {
			t.Errorf("got '%s', want  '%s'", got, want)
		}
	})
}

func TestGetToDoId(t *testing.T) {
	var buffer bytes.Buffer
	store := store.GetStore()
	store.Add("testToDo")
	t.Run("returns the id when a valid choice is made", func(t *testing.T) {

		got := GetToDoId(strings.NewReader("1\n"), &buffer)
		want := 1

		if got != want {
			t.Errorf("got %v but want %v", got, want)
		}

	})

	t.Run("prompts for valid number when unable to parse input", func(t *testing.T) {
		buffer.Reset()
		GetToDoId(strings.NewReader("cheese\n1\n"), &buffer)
		got := strings.TrimSpace(buffer.String())
		want := strings.TrimSpace(UnableToParseError)

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})

	t.Run("prompts for valid selection when unknown ID is entered", func(t *testing.T) {
		buffer.Reset()

		GetToDoId(strings.NewReader("2\n1\n"), &buffer)
		got := strings.TrimSpace(buffer.String())
		want := strings.TrimSpace(InvalidSelection)

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}

func TestGetToDo(t *testing.T) {

	t.Run("returns the correct to do", func(t *testing.T) {
		toDoId := 1
		got, _ := GetToDo(toDoId)

		if got.ID != toDoId {
			t.Errorf("got %v but want %v", got, toDoId)
		}
	})
	t.Run("returns error when to do with ID not found", func(t *testing.T) {
		_, err := GetToDo(2)
		want := errors.New(store.ToDoNotFoundError + " 2").Error()

		if err.Error() != want {
			t.Errorf("got %v but want %v", err.Error(), want)
		}
	})
}

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

func TestUpdateToDoStatus(t *testing.T) {
	t.Run("sets a to do to complete and back to incomplete", func(t *testing.T) {
		store := store.GetStore()
		store.ResetStore()
		store.Add("testThing")

		UpdateToDoStatus(strings.NewReader("1\n"), 1)
		got, _ := store.GetToDo(1)
		want := true

		if got.Completed != want {
			t.Errorf("got %v want %v", got.Completed, want)
		}

		UpdateToDoStatus(strings.NewReader("2\n"), 1)

		got, _ = store.GetToDo(1)
		want = false

		if got.Completed != want {
			t.Errorf("got %v want %v", got.Completed, want)
		}
	})

}

func TestDeleteToDo(t *testing.T) {
	store := store.GetStore()
	store.ResetStore()
	store.Add("testDelete")

	DeleteToDo(1)
	want := 0
	got := len(store.ToDos)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

var mockFunction = func(calledCount *int) func() {
	return func() {
		*calledCount++
	}
}
