package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type task struct {
	Description string
}

// create a slice to contain all of our tasks
var tasks []task

// create a function that
func addTask(description string) {
	task := task{Description: description}
	tasks = append(tasks, task)
	fmt.Println("task added ! ")
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("no tasks available")
		return
	}
	fmt.Println("These are the currents tasks...")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Description)
	}
}

func removeTasks(index int) {
	if index < 0 || index > len(tasks) {
		fmt.Println("invlid index value")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("task removed")

}

func main() {
	//creates a strcture to rep a task.
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nTodo List")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Remove Task")
		fmt.Println("4. Quit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("What is the task that you want to add?")
			scanner.Scan()
			description := strings.TrimSpace(scanner.Text())
			addTask(description)
			//hello
		}
	}
}
