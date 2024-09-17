package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	store "github.com/SamHenderson44/todo/internal/storePackage"
)

var inputChan = make(chan UserInput)
var printChan = make(chan string)
var printingCompleteCh = make(chan bool)
var mainMenuCh = make(chan bool)

var newToDoCh = make(chan bool)

type UserInput struct {
	inputType string
	input     string
}

func CliToDo() {

	go mainMenu()
	go printer()
	go HandleInput()
	go addNewToDoMenu()

	mainMenuCh <- true

	select {}
}

func HandleInput() {
	for msg := range inputChan {
		switch msg.inputType {
		case "mainMenu":
			mainMenuHandler(msg.input)
		case "addNewToDo":
			addNewToDo2(msg.input)
		default:
			mainMenuCh <- true
		}
	}

}

func mainMenuHandler(menuSelection string) {
	trimmed := strings.TrimSpace(menuSelection)

	switch trimmed {
	case "1":
		newToDoCh <- true
	case "2":
		store := store.GetStore()
		fmt.Println("---------------")
		fmt.Printf("%+v\n", store.GetToDos())
		fmt.Println("---------------")

		mainMenuCh <- true
	default:
		printChan <- "Bad input, try again"
		<-printingCompleteCh
		mainMenuCh <- true
	}
}

func addNewToDo2(newToDoTitle string) {
	store := store.GetStore()
	store.Add(newToDoTitle)
	printChan <- "Added new to do"
	<-printingCompleteCh
	mainMenuCh <- true
}

func mainMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		<-mainMenuCh
		fmt.Println(SelectAnOption)
		fmt.Println(AddNewTodo)
		fmt.Println(ShowCurrentToDos)
		fmt.Println(ChangerToDoStatus)
		fmt.Println(RemoveToDo)
		input, _ := reader.ReadString('\n')
		inputChan <- UserInput{inputType: "mainMenu", input: input}
	}
}

func addNewToDoMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		<-newToDoCh
		fmt.Println(NewToDoTextPrompt)
		input, _ := reader.ReadString('\n')
		inputChan <- UserInput{inputType: "addNewToDo", input: input}
	}
}

func printer() {
	for msg := range printChan {
		fmt.Println("------------")
		fmt.Println(msg)
		fmt.Println("------------")
		printingCompleteCh <- true
	}
}
