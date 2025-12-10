package main

import (
	"log"

	"github.com/Archiker-715/Task-Tracker/internal/task"
)

// todo: везде чекнуть file.Close

func main() {
	err := task.AddTask("buy car", "todo")
	if err != nil {
		log.Fatalf("add task error: %v", err)
	}

}
