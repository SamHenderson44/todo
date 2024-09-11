package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"text/tabwriter"
)

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
