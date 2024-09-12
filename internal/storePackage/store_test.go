package store

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("Adds correct title", func(t *testing.T) {
		newToDo := "new test to do"
		store := Store{}
		store.Add(newToDo)

		got := store.toDos[0].Title

		if got != newToDo {
			t.Errorf("got %s want %s", got, newToDo)
		}
	})

	t.Run("Adds correct title", func(t *testing.T) {
		newToDo := "new test to do"
		store := Store{}
		store.Add(newToDo)

		got := store.toDos[0].Title

		if got != newToDo {
			t.Errorf("got %s want %s", got, newToDo)
		}
	})

	t.Run("Sets the done status to 'false' for new to do", func(t *testing.T) {

		store := Store{}
		store.Add("test")
		want := false
		got := store.toDos[0].Completed

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

	})
}

func TestVGetToDos(t *testing.T) {
	want := ToDo{1, "test", false}
	store := Store{toDos: []ToDo{want}}
	got := store.GetToDos()[0]

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGetToDo(t *testing.T) {
	store := Store{toDos: []ToDo{
		{1, "test", false},
		{2, "test2", true},
	}}
	t.Run("Gets todo when given valid ID", func(t *testing.T) {
		got, _ := store.GetToDo(1)
		want := ToDo{1, "test", false}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}

	})
	t.Run("Returns error when given invalid ID", func(t *testing.T) {
		_, err := store.GetToDo(3)
		want := fmt.Errorf(ToDoNotFoundError + " 3")

		assertError(t, err, want)
	})
}

func TestDeleteToDo(t *testing.T) {
	store := Store{toDos: []ToDo{
		{1, "test", false},
		{2, "test2", true},
	}}
	t.Run("Deletes to do", func(t *testing.T) {
		store.DeleteToDo(1)
		_, err := store.GetToDo(1)
		want := fmt.Sprintf(ToDoNotFoundError + " 1")

		if err.Error() != want {
			t.Errorf("got %v, want %v", err.Error(), want)
		}

	})
	t.Run("Returns an error when given invalid id", func(t *testing.T) {
		err := store.DeleteToDo(3)
		want := fmt.Errorf(ToDoNotFoundError + " 3")

		assertError(t, err, want)
	})
}

func TestUpdateToDo(t *testing.T) {
	store := Store{toDos: []ToDo{
		{1, "test", false},
	}}
	t.Run("Updates to do", func(t *testing.T) {
		store.UpdateToDo(1, true)
		got := store.toDos[0].Completed
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Returns an error when given invalid id", func(t *testing.T) {
		err := store.UpdateToDo(3, true)
		want := fmt.Errorf(ToDoNotFoundError + " 3")

		assertError(t, err, want)
	})
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got.Error() != want.Error() {
		t.Errorf("got %v, want %v", got, want)
	}
}
