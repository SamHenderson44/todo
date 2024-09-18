package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/SamHenderson44/todo/internal/routes"
)

func main() {

	cliFlag := flag.Bool("cli", false, "Run cli")
	webFlag := flag.Bool("web", false, "Web app")
	cliFlag2 := flag.Bool("cli2", false, "Run cli with concurrency")

	flag.Parse()

	switch {
	case *cliFlag:
		ShowToDoOptions(os.Stdin)
	case *webFlag:
		routes.InitRoutes()
	case *cliFlag2:
		CliToDo()
	default:
		fmt.Println("Please provide a valid flag. Use --help to see available options.")
	}

}
