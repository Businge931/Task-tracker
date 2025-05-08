package main

import (
	"fmt"
	"os"

	"github.com/Businge931/tasktracker/internal/cli"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		os.Exit(1)
	}

	var err error
	switch args[1] {
	case "add":
		err = cli.HandleAdd(args)
	case "update":
		err = cli.HandleUpdate(args)
	case "delete":
		err = cli.HandleDelete(args)
	case "mark":
		err = cli.HandleMark(args)
	case "list":
		err = cli.HandleList(args)
	default:
		err = cli.HandleDynamicMark(args)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
