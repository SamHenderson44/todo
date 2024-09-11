package main

import (
	"fmt"
)

var data = 0

func main() {
	go odds()
	go evens()

}

func odds() {
	for i := 1; i < 10000; i += 2 {

		data = i
		fmt.Printf("odds %v: \n", i)

	}
}

func evens() {
	for i := 0; i < 10000; i += 2 {

		data = i
		fmt.Printf("even %v: \n", i)

	}
}
