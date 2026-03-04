package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Completed   bool      `json:"completed"`
}

const file = "tasks.json"

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(file)
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

func saveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0644)
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

func relativeTime(t time.Time) string {
	duration := time.Since(t)
	hours, minutes := duration.Hours(), duration.Minutes()
	if hours >= 24 {
		numDays := int(hours / 24)
		var article string
		if numDays > 1 {
			article = "days"
		} else {
			article = "day"
		}
		return fmt.Sprintf("%d %s ago", numDays, article)
	}
	if minutes >= 60 {
		numHours := int(minutes / 60)
		var article string
		if numHours > 1 {
			article = "hours"
		} else {
			article = "hour"
		}
		return fmt.Sprintf("%d %s ago", numHours, article)
	}
	if int(minutes) >= 1 {
		var article string
		if minutes > 1 {
			article = "minutes"
		} else {
			article = "minute"
		}
		return fmt.Sprintf("%d %s ago", int(minutes), article)
	}
	return "a few seconds ago"
}

func listTasks(tasks []Task, all bool) {
	w := tabwriter.NewWriter(os.Stdout, 10, 0, 5, ' ', 0)
	if all {
		fmt.Fprintln(w, "ID\tDescription\tCreated\tDone")
		for _, val := range tasks {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", val.ID, val.Description, relativeTime(val.CreatedAt), val.Completed)
		}
	} else {
		fmt.Fprintln(w, "ID\tDescription\tCreated")
		for _, val := range tasks {
			if !val.Completed {
				fmt.Fprintf(w, "%v\t%v\t%v\n", val.ID, val.Description, relativeTime(val.CreatedAt))
			}
		}
	}
	w.Flush()
}

func completeTask(tasks []Task, id int) error {
	var taskIdx int
	found := false
	for idx, val := range tasks {
		if val.ID == id {
			taskIdx = idx
			found = true
		}
	}
	if !found {
		return errors.New("Error: Task not found")
	}
	tasks[taskIdx].Completed = true
	return nil
}

func deleteTask(tasks []Task, id int) ([]Task, error) {
	var taskIdx int
	found := false
	for idx, val := range tasks {
		if val.ID == id {
			taskIdx = idx
			found = true
		}
	}
	if !found {
		return nil, errors.New("Error: Task not found")
	}
	result := make([]Task, 0, len(tasks)-1)
	result = append(result, tasks[:taskIdx]...)
	result = append(result, tasks[taskIdx+1:]...)
	return result, nil
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
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		tasks = addTask(tasks, d)
		err = saveTasks(tasks)
		if err != nil {
			fmt.Println("Error: Failed to save task")
			os.Exit(1)
		}
		fmt.Println("Successfully saved task")
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listAll := listCmd.Bool("all", false, "List all tasks including completed")
		listCmd.Parse(os.Args[2:])
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		listTasks(tasks, *listAll)
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
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		err = completeTask(tasks, id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = saveTasks(tasks)
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
		tasks, err := loadTasks()
		if err != nil {
			fmt.Println("Error: Failed to load tasks")
			os.Exit(1)
		}
		updatedTasks, err := deleteTask(tasks, id)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = saveTasks(updatedTasks)
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
