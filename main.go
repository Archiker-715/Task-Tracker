package main

import (
	"log"

	"github.com/Archiker-715/Task-Tracker/internal/task"
)

var tasks *task.Tasks = &task.Tasks{}

func main() {

	// addTask()

	updateTask(1)
}

func addTask() {
	err := tasks.AddTask("buy car")
	if err != nil {
		log.Fatalf("add task error: %v", err)
	}
}

func updateTask(taskId int) {
	err := tasks.UpdateTask(taskId, "b")
	if err != nil {
		log.Fatalf("update task error: %v", err)
	}
}
