package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

const (
	NewToDoTextPrompt  = "What do you want to do?"
	SelectAnOption     = "Select an option:"
	AddNewTodo         = "1. Add new to do"
	ShowCurrentToDos   = "2. Show to dos"
	ChangerToDoStatus  = "3. Change to do status"
	InvalidSelection   = "Invalid choice, please try again."
	ChooseToDo         = "Choose a to do to update"
	ChangeStatusPrompt = "what do you want to change the status to?"
	CompleteStatus     = "1. Complete"
	IncompleteStatus   = "2. Incomplete"
	UnableToParseError = "Please enter a valid number"
	CompleteToDo       = true
	IncompleteToDo     = false
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

var GetToDo = func(ID int) (store.ToDo, error) {
	todo, err := HandleGetToDo(ID)

	if err != nil {
		return todo, err
	}
	return todo, nil
}

func GetToDoId(reader io.Reader, writer io.Writer) int {
	ShowToDos()
	scanner := bufio.NewScanner(reader)
	for {
		fmt.Println(ChooseToDo)
		scanner.Scan()
		selection := strings.TrimSpace(scanner.Text())
		toDoId, parseErr := strconv.Atoi(selection)
		if parseErr != nil {
			fmt.Fprintln(writer, UnableToParseError)
			continue
		}
		_, err := GetToDo(toDoId)
		if err != nil {
			fmt.Fprintln(writer, InvalidSelection)
		} else {
			return toDoId
		}
	}
}

func SelectNewToDoStatus(reader io.Reader) string {
	scanner := bufio.NewScanner(reader)
	fmt.Println(ChangeStatusPrompt)
	fmt.Println(CompleteStatus)
	fmt.Println(IncompleteStatus)

	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

var UpdateToDoStatus = func(reader io.Reader, toDoId int) {

	for {
		selection := SelectNewToDoStatus(reader)
		switch selection {
		case "1":
			HandleUpdateToDo(toDoId, CompleteToDo)
			return
		case "2":
			HandleUpdateToDo(toDoId, IncompleteToDo)
			return
		default:
			InvalidInput()
		}

	}

}

func ShowToDoOptions(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for {

		fmt.Println(SelectAnOption)
		fmt.Println(AddNewTodo)
		fmt.Println(ShowCurrentToDos)
		fmt.Println(ChangerToDoStatus)

		if !scanner.Scan() {
			break
		}
		selection := strings.TrimSpace(scanner.Text())

		switch selection {
		case "1":
			ReadToDoText()
		case "2":
			ShowToDos()
		case "3":
			toDoId := GetToDoId(reader, os.Stdout)
			UpdateToDoStatus(os.Stdin, toDoId)
		default:
			InvalidInput()
		}
	}
}
