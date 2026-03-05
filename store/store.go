package store

import (
	"encoding/json"
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

func LoadTasks() ([]Task, error) {
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

func SaveTasks(tasks []Task) error {
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
