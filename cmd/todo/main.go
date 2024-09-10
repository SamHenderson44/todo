package main

import (
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

func main() {
	toDos := []ToDoItem{
		{Item: "ToDo 1", Completed: true},
		{Item: "ToDo 2", Completed: false},
		{Item: "ToDo 3", Completed: true},
		{Item: "ToDo 4", Completed: true},
		{Item: "ToDo 5", Completed: false},
		{Item: "ToDo 6", Completed: true},
		{Item: "ToDo 7", Completed: false},
		{Item: "ToDo 8", Completed: false},
		{Item: "ToDo 9", Completed: true},
		{Item: "ToDo 10", Completed: true},
	}
	PrintToDos(os.Stdout, toDos...)
}

type ToDoItem struct {
	Item      string
	Completed bool
}

func PrintToDos(writer io.Writer, toDos ...ToDoItem) {
	tabWriter := tabwriter.NewWriter(writer, 20, 0, 1, ' ', 0)
	fmt.Fprintln(tabWriter, "To Do\tCompleted")
	for _, todo := range toDos {
		fmt.Fprintf(tabWriter, "%v\t%v\n", todo.Item, todo.Completed)
	}
	tabWriter.Flush()
}
