package main

import (
	"fmt"
	"sort"
)

// Task struct to hold task details
type Task struct {
	ID          int
	Description string
	Priority    int
}

// Global slice to store tasks
var tasks []Task

// Global task ID counter
var nextID int = 1

// Function to add a task
func addTask(description string, priority int) {
	if description == "" {
		fmt.Println("Error: Task description cannot be empty.")
		return
	}
	if priority < 1 || priority > 3 {
		fmt.Println("Error: Priority must be between 1 (High) and 3 (Low).")
		return
	}

	task := Task{
		ID:          nextID,
		Description: description,
		Priority:    priority,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added successfully!")
}

// Function to view tasks sorted by priority
func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to display.")
		return
	}

	// Sort tasks by priority (High to Low)
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})

	fmt.Println("\nTasks:")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Priority: %d\n", task.ID, task.Description, task.Priority)
	}
	fmt.Println()
}

// Function to delete a task by ID
func deleteTask(id int) {
	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Error: Task not found.")
		return
	}

	// Remove task from slice
	tasks = append(tasks[:index], tasks[index+1:]...)
	fmt.Println("Task deleted successfully!")
}

// Function to update a task by ID
func updateTask(id int, newDescription string, newPriority int) {
	for i, task := range tasks {
		if task.ID == id {
			if newDescription != "" {
				tasks[i].Description = newDescription
			}
			if newPriority >= 1 && newPriority <= 3 {
				tasks[i].Priority = newPriority
			} else {
				fmt.Println("Error: Priority must be between 1 (High) and 3 (Low).")
				return
			}
			fmt.Println("Task updated successfully!")
			return
		}
	}

	fmt.Println("Error: Task not found.")
}

// Function to display menu and handle user input
func showMenu() {
	for {
		fmt.Println("\nTo-Do List Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Update Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var description string
			var priority int
			fmt.Print("Enter task description: ")
			fmt.Scanln(&description)
			fmt.Print("Enter task priority (1 = High, 2 = Medium, 3 = Low): ")
			fmt.Scan(&priority)
			addTask(description, priority)
		case 2:
			viewTasks()
		case 3:
			var id int
			fmt.Print("Enter task ID to delete: ")
			fmt.Scan(&id)
			deleteTask(id)
		case 4:
			var id int
			var newDescription string
			var newPriority int
			fmt.Print("Enter task ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter new task description (leave empty to keep unchanged): ")
			fmt.Scanln(&newDescription)
			fmt.Print("Enter new priority (1 = High, 2 = Medium, 3 = Low, or 0 to keep unchanged): ")
			fmt.Scan(&newPriority)
			updateTask(id, newDescription, newPriority)
		case 5:
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func main() {
	showMenu()
}




