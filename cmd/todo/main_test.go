package main

import (
	"bytes"
	"testing"
)

func TestPrintToDos(t *testing.T) {

	toDos := []ToDoItem{
		{Item: "Task 1", Completed: true},
		{Item: "Task 2", Completed: false},
	}
	var buffer bytes.Buffer

	t.Run("Prints the tab headings", func(t *testing.T) {
		buffer.Reset()
		PrintToDos(&buffer)
		expected := "To Do               Completed\n"
		if buffer.String() != expected {
			t.Errorf("Expected output %q, but got %q", expected, buffer.String())
		}

	})

	t.Run("Prints to do items correctly", func(t *testing.T) {
		buffer.Reset()
		PrintToDos(&buffer, toDos...)
		expected := "To Do               Completed\nTask 1              true\nTask 2              false\n"
		if buffer.String() != expected {
			t.Errorf("Expected output %q, but got %q", expected, buffer.String())
		}

	})
}
