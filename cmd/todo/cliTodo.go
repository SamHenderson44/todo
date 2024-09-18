package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

const (
	MainMenuConst         = "MainMenu"
	AddNewToDoConst       = "addNewToDo"
	UpdateToDoConst       = "updateToDo"
	UpdateHandlerError    = "Error inside update Handler"
	ChangeToDoStatusError = "Error inside changeToDoStatus"
	BadInputError         = "Bad input, try again"
	ChooseToDoStatus      = "Choose a status\n\n1. Complete\n2. Incomplete"
	ChooseToDoToUpdate    = "Choose a To Do to update:\n\n"
	AddedNewToDo          = "Added new to do"
	LineBreak             = "------------"
)

var mainMenuCh = make(chan bool)
var inputCh = make(chan UserInput)
var updateToDoCh = make(chan int)
var printCh = make(chan string)
var printingCompleteCh = make(chan bool)

var newToDoCh = make(chan bool)

type UserInput struct {
	inputType string
	input     string
}

func CliToDo() {

	go mainMenu()
	go handleInput()
	go addNewToDoMenu()
	go changeToDoStatus()
	go printer()

	mainMenuCh <- true

	select {}
}

func mainMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		<-mainMenuCh
		fmt.Println(SelectAnOption)
		fmt.Println(AddNewTodo)
		fmt.Println(ShowCurrentToDos)
		fmt.Println(ChangerToDoStatus)
		input, _ := reader.ReadString('\n')
		trimmedInput := strings.TrimSpace(input)
		inputCh <- UserInput{inputType: MainMenuConst, input: trimmedInput}
	}
}

func mainMenuHandler(menuSelection string) {
	trimmed := strings.TrimSpace(menuSelection)
	toDos := store.GetStore().GetToDos()

	switch trimmed {
	case "1":
		newToDoCh <- true
	case "2":
		formatted := store.FormatToDos(toDos)
		printCh <- formatted
		<-printingCompleteCh
		mainMenuCh <- true
	case "3":
		// Ran into problems here trying to use a channel to print out the "choose a todo to update" menu and
		// the subsequent "choose a status to update for the chosen to to".
		updateHandler()
	default:
		printCh <- BadInputError
		<-printingCompleteCh
		mainMenuCh <- true
	}
}

func handleInput() {
	for msg := range inputCh {
		switch msg.inputType {
		case MainMenuConst:
			mainMenuHandler(msg.input)
		case AddNewToDoConst:
			addNewToDo(msg.input)
		case UpdateToDoConst:
			updateHandler()
		default:
			mainMenuCh <- true
		}
	}
}

func addNewToDoMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		<-newToDoCh
		fmt.Println(NewToDoTextPrompt)
		input, _ := reader.ReadString('\n')
		inputCh <- UserInput{inputType: AddNewToDoConst, input: input}
	}
}

func addNewToDo(newToDoTitle string) {
	store := store.GetStore()
	store.Add(newToDoTitle)
	printCh <- AddedNewToDo
	<-printingCompleteCh
	mainMenuCh <- true
}

func changeToDoStatus() {
	store := store.GetStore()
	reader := bufio.NewReader(os.Stdin)

	for msg := range updateToDoCh {
		printCh <- ChooseToDoStatus
		<-printingCompleteCh

		input, _ := reader.ReadString('\n')
		inputInt, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println(ChangeToDoStatusError)
		}

		status := inputInt == 1
		store.UpdateToDo(msg, status)
		mainMenuCh <- true
	}
}

func updateHandler() {

	toDos := store.GetStore().GetToDos()
	formattedToDos := store.FormatToDos(toDos)

	printCh <- fmt.Sprintf(ChooseToDoToUpdate+"%s", formattedToDos)
	<-printingCompleteCh

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	toDoId, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(UpdateHandlerError)
	}
	updateToDoCh <- toDoId

}

func printer() {
	for msg := range printCh {
		fmt.Println("\n" + LineBreak)
		fmt.Println(msg)
		fmt.Println(LineBreak)
		printingCompleteCh <- true
	}
}
