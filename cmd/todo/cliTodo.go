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
	MainMenuConst   = "MainMenu"
	AddNewToDoConst = "addNewToDo"
	UpdateToDoConst = "updateToDo"
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
		trimmed := strings.TrimSpace(input)
		inputCh <- UserInput{inputType: MainMenuConst, input: trimmed}
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
		updateHandler()
	default:
		printCh <- "Bad input, try again"
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
	printCh <- "Added new to do"
	<-printingCompleteCh
	mainMenuCh <- true
}

func changeToDoStatus() {
	store := store.GetStore()
	reader := bufio.NewReader(os.Stdin)

	for msg := range updateToDoCh {
		printCh <- "Choose a status\n\n1. Complete\n2. Incomplete"
		<-printingCompleteCh

		input, _ := reader.ReadString('\n')
		inputInt, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Error inside changeToDoStatus")
		}

		status := inputInt == 1
		store.UpdateToDo(msg, status)
		mainMenuCh <- true
	}
}

func updateHandler() {

	toDos := store.GetStore().GetToDos()
	formatted := store.FormatToDos(toDos)

	printCh <- fmt.Sprintf("Choose a To Do to update:\n\n%s", formatted)
	<-printingCompleteCh

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	toDoId, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Error inside update Handler")
	}
	updateToDoCh <- toDoId

}

func printer() {
	for msg := range printCh {
		fmt.Println("\n------------")
		fmt.Println(msg)
		fmt.Println("------------")
		printingCompleteCh <- true
	}
}
