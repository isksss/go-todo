package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/isksss/go-todo"
)

var (
	reader *bufio.Reader = bufio.NewReader(os.Stdin)
)

func main() {
	fmt.Println("Welcome to ToDo Manager CLI!")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("Exiting ToDo Manager CLI...")
			break
		}
		switch input {
		case "help":
			fmt.Println("List of commands:")
			fmt.Println("  help    : display list of commands")
			fmt.Println("  add     : add new task")
			fmt.Println("  list    : list all tasks")
			fmt.Println("  done    : mark a task as done")
			fmt.Println("  delete  : delete a task")
			fmt.Println("  exit    : exit ToDo Manager CLI")
		case "add":
			fmt.Println("Adding new task...")
			// TODO: add implementation for adding task
			addCommand()
		case "list":
			fmt.Println("Listing all tasks...")
			// TODO: add implementation for listing tasks
		case "done":
			fmt.Println("Marking a task as done...")
			// TODO: add implementation for marking task as done
		case "delete":
			fmt.Println("Deleting a task...")
			// TODO: add implementation for deleting task
		default:
			fmt.Printf("Invalid command: %s\n", input)
		}
	}
}

func addCommand() {
	var t todo.ToDo

	fmt.Print("Name >>> ")
	input, _ := reader.ReadString('\n')
	t.Name = strings.TrimSpace(input)
	t.Done = false

	t.AddTask()
}
