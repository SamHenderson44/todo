package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"

	fileio "github.com/SamHenderson44/todo/internal/fileIoPackage"
)

var toDoItems = []ToDoItem{
	{Item: "ToDo 1", Completed: true},
	{Item: "ToDo 2", Completed: false},
	{Item: "ToDo 3", Completed: true},
	{Item: "ToDo 4", Completed: true},
	{Item: "ToDo 5", Completed: false},
	{Item: "ToDo 6", Completed: true},
	{Item: "ToDo 7", Completed: false},
	{Item: "ToDo 8", Completed: false},
	{Item: "ToDo 9", Completed: true},
	{Item: "ToDo 12", Completed: true},
}

func main() {
	dummyFileName := "dummy.json"
	PrintToDos(os.Stdout, toDoItems...)
	jsonToDos, _ := CreateJsonToDos(toDoItems...)
	file, _ := fileio.CreateFile(dummyFileName)
	fileio.WriteToFile(file, jsonToDos)

	file, err := os.Open(dummyFileName)
	if err != nil {
		fmt.Println("error")
	}
	openedFile, err := fileio.ReadFile(file)
	if err != nil {
		fmt.Println("error")
	}

	PrintJsonToDos(openedFile)

}

type ToDoItem struct {
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

func PrintToDos(writer io.Writer, toDos ...ToDoItem) {
	tabWriter := tabwriter.NewWriter(writer, 20, 0, 1, ' ', 0)
	fmt.Fprintln(tabWriter, "To Do\tCompleted")
	for _, todo := range toDos {
		fmt.Fprintf(tabWriter, "%v\t%v\n", todo.Item, todo.Completed)
	}
	tabWriter.Flush()
}

func CreateJsonToDos(toDos ...ToDoItem) ([]byte, error) {
	jsonToToDos, err := json.Marshal(toDos)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	return jsonToToDos, err
}

func PrintJsonToDos(toDos []byte) {
	var unmarshalledToDos []ToDoItem
	err := json.Unmarshal(toDos, &unmarshalledToDos)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	formattedJson, _ := json.MarshalIndent(unmarshalledToDos, "", "  ")

	fmt.Println(string(formattedJson))
}
