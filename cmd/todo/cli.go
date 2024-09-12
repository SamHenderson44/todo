package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	NewToDoTextPrompt = "What do you want to do?"
	SelectAnOption    = "Select an option:"
	AddNewTodo        = "1. Add new to do"
	ShowCurrentToDos  = "2. Show to dos"
	InvalidSelection  = "Invalid choice, please try again."
)

var ReadToDoText = func() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(NewToDoTextPrompt)

	scanner.Scan()
	toDoText := strings.TrimSpace(scanner.Text())
	HandleSaveNewToDo(toDoText)
}

var ShowToDos = func() {
	toDos := HandleGetToDos()
	PrintToDos(os.Stdout, toDos...)
}

var InvalidInput = func() {
	fmt.Println(InvalidSelection)
}

func ShowToDoOptions(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for {

		fmt.Println(SelectAnOption)
		fmt.Println(AddNewTodo)
		fmt.Println(ShowCurrentToDos)

		if !scanner.Scan() {
			break
		}
		selection := strings.TrimSpace(scanner.Text())

		switch selection {
		case "1":
			ReadToDoText()
		case "2":
			ShowToDos()
		default:
			InvalidInput()
		}
	}
}
