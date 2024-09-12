package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	NewToDoTextPrompt = "\nWhat do you want to do?"
	SelectAnOption    = "\nSelect an option:"
	AddNewTodo        = "1. Add new to do"
	InvalidSelection  = "\nInvalid choice, please try again."
)

func ReadToDoText() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(NewToDoTextPrompt)

	scanner.Scan()
	toDoText := strings.TrimSpace(scanner.Text())
	HandleSaveNewToDo(toDoText)
	ShowToDoOptions(os.Stdin)
}

func ShowToDoOptions(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for {

		fmt.Println(SelectAnOption)
		fmt.Println(AddNewTodo)

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			ReadToDoText()
			return
		default:
			fmt.Println(InvalidSelection)
		}
	}
}
