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




