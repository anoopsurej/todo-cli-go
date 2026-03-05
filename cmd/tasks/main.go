package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	logic "github.com/anoopsurej/todo-cli-go/internal/tasks"
	"github.com/anoopsurej/todo-cli-go/store"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Undefined description")
			os.Exit(1)
		}
		d := os.Args[2]
		tasks, err := store.LoadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		tasks = logic.AddTask(tasks, d)
		err = store.SaveTasks(tasks)
		if err != nil {
			fmt.Println("Error: Failed to save task")
			os.Exit(1)
		}
		fmt.Println("Successfully saved task")
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		var listAll bool
		listCmd.BoolVar(&listAll, "all", false, "List all tasks including completed")
		listCmd.BoolVar(&listAll, "a", false, "List all tasks including completed")
		listCmd.Parse(os.Args[2:])
		tasks, err := store.LoadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		logic.ListTasks(tasks, listAll)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID not provided")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID")
			os.Exit(1)
		}
		tasks, err := store.LoadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		err = logic.CompleteTask(tasks, id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = store.SaveTasks(tasks)
		if err != nil {
			fmt.Println("Error: Failed to save task")
			os.Exit(1)
		}
		fmt.Println("Successfully completed task")
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: ID not provided")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid ID")
			os.Exit(1)
		}
		tasks, err := store.LoadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		updatedTasks, err := logic.DeleteTask(tasks, id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = store.SaveTasks(updatedTasks)
		if err != nil {
			fmt.Println("Error: Failed to save task")
			os.Exit(1)
		}
		fmt.Println("Successfully deleted task")
	default:
		fmt.Println("Error: Urecognized command")
		os.Exit(1)
	}
}
