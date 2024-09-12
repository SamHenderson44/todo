package main

import (
	"bytes"
	"testing"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

var toDos = []store.ToDo{
	{Title: "Task 1", Completed: true},
	{Title: "Task 2", Completed: false},
}

func TestPrintToDos(t *testing.T) {
	var buffer bytes.Buffer

	t.Run("Prints the tab headings", func(t *testing.T) {
		buffer.Reset()
		PrintToDos(&buffer)
		expected := "ID                  To Do               Completed\n"
		if buffer.String() != expected {
			t.Errorf("Expected output %q, but got %q", expected, buffer.String())
		}

	})

	t.Run("Prints to do items correctly", func(t *testing.T) {
		buffer.Reset()
		PrintToDos(&buffer, toDos...)
		// expected := "To Do               Completed\nTask 1              true\nTask 2              false\n"
		expected := "ID                  To Do               Completed\n0                   Task 1              true\n0                   Task 2              false\n"
		if buffer.String() != expected {
			t.Errorf("Expected output %q, but got %q", expected, buffer.String())
		}

	})
}

func TestCreateJsonToDos(t *testing.T) {
	t.Run("successfully marshals to do", func(t *testing.T) {
		_, err := CreateJsonToDos(toDos...)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

}
