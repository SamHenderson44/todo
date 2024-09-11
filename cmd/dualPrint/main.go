package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

type ToDo struct {
	item   string
	status bool
}

var toDos = []ToDo{
	{"Take out trash", false},
	{"spend money", true},
	{"go to sleep", true},
}

func printToDoItem(itemCh, statusCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, p := range toDos {
		<-itemCh
		logger(os.Stdout, p.item)
		statusCh <- true
	}

}

func printToDoStatus(itemCh, statusCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i, p := range toDos {
		<-statusCh
		logger(os.Stdout, strconv.FormatBool(p.status))
		if i < len(toDos) {
			itemCh <- true
		}
	}
}

func logger(writer io.Writer, item string) {
	fmt.Fprintln(writer, item)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	itemCh := make(chan bool, 1)
	statusCh := make(chan bool, 1)

	go printToDoItem(itemCh, statusCh, &wg)
	go printToDoStatus(itemCh, statusCh, &wg)

	itemCh <- true

	wg.Wait()
}
