package task

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func fillTask(file *os.File, taskDescription, taskStatus string, fileSize int) {

	allTasks := &Tasks{}

	if fileSize > 0 {
		if err := json.Unmarshal(readFile(file), allTasks); err != nil {
			log.Fatalf("unmarshalling err: %v", err)
		}
	}

	maxId := getMaxId(*allTasks)

	newTask := Task{
		Id:          maxId + 1,
		Description: taskDescription,
		Status:      taskStatus,
		CreatedAt:   time.Now(),
	}

	allTasks.Tasks = append(allTasks.Tasks, newTask)

	b, err := json.Marshal(allTasks)
	if err != nil {
		log.Fatalf("marshalling err: %v", err)
	}

	seekPosition(file)

	if _, err := file.Write(b); err != nil {
		log.Fatalf("fill taskfile error: %v", err)
	}
}

func getMaxId(allTasks Tasks) int {
	var maxId int

	if len(allTasks.Tasks) > 0 {
		maxId = allTasks.Tasks[0].Id
		for _, task := range allTasks.Tasks {
			if task.Id > maxId {
				maxId = task.Id
			}
		}
		return maxId
	}
	return 0
}
