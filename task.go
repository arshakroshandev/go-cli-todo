package main

import "fmt"

type Task struct {
	Name string
	Done bool
}

var tasks []Task

func addTask(name string) {
	task := Task{
		Name: name,
		Done: false,
	}
	tasks = append(tasks, task)
	saveTasks()
	fmt.Println("Task added")
}

func markTaskDone(index int) {
	if index < 0 || index > len(tasks) {
		fmt.Println("Invalid task number")
		return
	}
	tasks[index].Done = true
	saveTasks()
	fmt.Println("Task marked complete")
	return
}

func deleteTask(index int) {
	if index < 0 || index > len(tasks) {
		fmt.Println("Invalid task number")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTasks()
	fmt.Println("Task deleted")
	return
}

func listTask() {
	if len(tasks) == 0 {
		fmt.Println("No task found, please add tasks")
		return
	}
	fmt.Println("\nTasks:")
	for i, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Name, status)
	}
	return
}
