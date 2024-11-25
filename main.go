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






