package main

import (
	"bufio" // buffer input / output
	"fmt"
	"os"
	"strconv" // string to int
	"strings" // string manipulation
)

type Task struct {
	Text string
	Completed bool
}

func main() {
	tasks := []Task{}

	for {
		showMenu()
		option := getUserInput("Enter your choices")

		switch option {
		case "1":
			showTasks(tasks)
		case "2":
			addTasks(&tasks)
		case "3":
			markTaskComplete(&tasks)
		case "4":
			saveTaskToFile(tasks)
		case "5":
			fmt.Println("Exiting to-do app.")
			return
		default:
			fmt.Println("Invalid choice. Please try again")
		}
	}
}

func showMenu() {
	fmt.Println("\nMenu")
	fmt.Println("1. Show Tasks")
	fmt.Println("2. Add Tasks")
	fmt.Println("3. Mark Task as Complete")
	fmt.Println("4. Save Tasks to File")
	fmt.Println("5. Exit")
}

func getUserInput(prompt string) string { // takes a prompt as input and return as a cleaned string
	reader := bufio.NewReader(os.Stdin) // reads keyboard input
	fmt.Println(prompt) 
	input, _ := reader.ReadString('\n') // read the input until they press Enter
	return strings.TrimSpace(input) //use strings package to remove leading/trailing spaces
}

func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("You have no tasks.")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
	}
}

func addTasks(tasks *[]Task) {
	taskText := getUserInput("Enter Task Description")
	*tasks = append(*tasks, Task{Text: taskText}) // append a new task to 'tasks' slice, use a pointer to the slice (*task)
	fmt.Println("Task has been added.")
}

func markTaskComplete(tasks *[]Task) {
	showTasks(*tasks)
	taskIndexStr := getUserInput("Enter Task number to mark as completed: ")
	taskIndex, err := strconv.Atoi(taskIndexStr) // Convert string to int
	if err != nil || taskIndex < 1 || taskIndex > len(*tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	(*tasks)[taskIndex - 1].Completed = true
	fmt.Println("Task marked as completed.")
}

func saveTaskToFile(tasks []Task) {
	file, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "x"
		}
		file.WriteString(fmt.Sprintf("[%s] %s\n", status, task.Text))
	}
	fmt.Println("Your tasks have been saved to file named 'tasks.txt'")
}

