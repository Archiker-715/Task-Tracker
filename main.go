package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Archiker-715/Task-Tracker/constants"
	"github.com/Archiker-715/Task-Tracker/internal/task"
)

var tasks *task.Tasks = &task.Tasks{}

func main() {

	var (
		taskDecription string
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Hello! Insert your operation with task\n")
	usersInput, _ := reader.ReadString('\n')
	formattedInput := strings.TrimSpace(usersInput)
	inp := strings.Split(formattedInput, " ")

	if strings.Contains(formattedInput, `"`) {
		firstIdx := strings.Index(formattedInput, `"`)
		if firstIdx == -1 {
			log.Println("not found")
		}
		firstPartString := formattedInput[firstIdx+1:]

		secondIdx := strings.Index(firstPartString, `"`)
		if firstIdx == -1 {
			log.Println("not found")
		}
		taskDecription = firstPartString[:secondIdx]
	}

	switch inp[0] {
	case constants.List:
		if len(inp) == 1 {
			listTasks()
		} else {
			switch inp[1] {
			case constants.Done:
				filteredListTasks(inp[1])
			case constants.Todo:
				filteredListTasks(inp[1])
			case constants.InProgress:
				filteredListTasks(inp[1])
			default:
				// err
			}
		}
	case constants.Add:
		addTask(taskDecription)
		fmt.Printf("Task added successfully (ID: \n")
	case constants.MarkInProgress:
		taskId, err := strconv.Atoi(inp[1])
		if err != nil {
			// err
		}
		updateTask(taskId, inp[0])
	case constants.MarkDone:
		taskId, err := strconv.Atoi(inp[1])
		if err != nil {
			// err
		}
		updateTask(taskId, inp[0])
	case constants.Delete:
		taskId, err := strconv.Atoi(inp[1])
		if err != nil {
			// err
		}
		deleteTask(taskId)
	case constants.Update:
		taskId, err := strconv.Atoi(inp[1])
		if err != nil {
			// err
		}
		updateTask(taskId, taskDecription)
	default:
		// err
	}

}

func addTask(taskDescription string) {
	err := tasks.AddTask(taskDescription)
	if err != nil {
		log.Fatalf("add task error: %v", err)
	}
}

func listTasks() {
	err := tasks.ListTasks()
	if err != nil {
		log.Fatalf("add task error: %v", err)
	}
}

func filteredListTasks(filter string) {
	err := tasks.FilteredListTasks(filter)
	if err != nil {
		log.Fatalf("add task error: %v", err)
	}
}

func updateTask(taskId int, taskData string) {
	err := tasks.UpdateTask(taskId, taskData)
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
