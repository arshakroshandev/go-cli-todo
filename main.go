package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	loadTasks()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n ---TO DO APP---")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Task")
		fmt.Println("3. Mark Task Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Println("Please choose an option")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {
		case 1:
			fmt.Println("Enter task name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			addTask(name)

		case 2:
			listTask()
		case 3:
			fmt.Println("Enter Task Number")
			input, _ := reader.ReadString('\n')
			number, _ := strconv.Atoi(strings.TrimSpace(input))
			markTaskDone(number - 1)
		case 4:
			fmt.Println("Enter Task Number")
			input, _ := reader.ReadString('\n')
			number, _ := strconv.Atoi(strings.TrimSpace(input))
			deleteTask(number - 1)
		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
