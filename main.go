package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
}

const file = "tasks.json"

func loadTasks(filename string) ([]Task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return []Task{}, err
	}
	var v []Task
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return v, nil
}

func saveTasks(filename string, tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func addTask(tasks []Task, description string) []Task {
	var id int
	if len(tasks) == 0 {
		id = 1
	} else {
		id = tasks[len(tasks)-1].ID + 1
	}
	t := Task{
		ID:          id,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	}

	return append(tasks, t)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Not enough arguments")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Undefined description")
			os.Exit(1)
		}
		d := os.Args[2]
		tasks, err := loadTasks(file)
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		tasks = addTask(tasks, d)
		err = saveTasks(file, tasks)
		if err != nil {
			fmt.Println("Error: Failed to save task")
			os.Exit(1)
		}
		fmt.Println("Successfully saved task")
		os.Exit(0)
	default:
		fmt.Println("Error: Urecognized command")
		os.Exit(1)
	}
}
