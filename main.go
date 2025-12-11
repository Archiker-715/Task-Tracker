package main

import (
	"log"

	"github.com/Archiker-715/Task-Tracker/internal/task"
)

var tasks *task.Tasks = &task.Tasks{}

func main() {

	// addTask("buy car 1")
	// addTask("buy car 2")

	// updateTask(1)

	// deleteTask(2)
}

func addTask(taskDescription string) {
	err := tasks.AddTask(taskDescription)
	if err != nil {
		log.Fatalf("add task error: %v", err)
	}
}

func updateTask(taskId int) {
	err := tasks.UpdateTask(taskId, "butttttt")
	if err != nil {
		log.Fatalf("update task error: %v", err)
	}
}

func deleteTask(taskId int) {
	err := tasks.DeleteTask(taskId)
	if err != nil {
		log.Fatalf("delete task error: %v", err)
	}
}
