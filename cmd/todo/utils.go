package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"text/tabwriter"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

func PrintToDos(writer io.Writer, toDos ...store.ToDo) {
	tabWriter := tabwriter.NewWriter(writer, 20, 0, 1, ' ', 0)
	fmt.Fprintln(tabWriter, "ID\tTo Do\tCompleted")
	for _, todo := range toDos {
		fmt.Fprintf(tabWriter, "%v\t%v\t%v\n", todo.ID, todo.Title, todo.Completed)
	}
	tabWriter.Flush()
}

func CreateJsonToDos(toDos ...store.ToDo) ([]byte, error) {
	jsonToToDos, err := json.Marshal(toDos)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}
	return jsonToToDos, err
}

func PrintJsonToDos(toDos []byte) {
	var unmarshalledToDos []store.ToDo
	err := json.Unmarshal(toDos, &unmarshalledToDos)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	formattedJson, _ := json.MarshalIndent(unmarshalledToDos, "", "  ")

	fmt.Println(string(formattedJson))
}
