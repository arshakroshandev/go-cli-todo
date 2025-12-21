package main

import (
	"encoding/json"
	"os"
)

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func loadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil
	}
	return json.Unmarshal(data, &tasks)
}
