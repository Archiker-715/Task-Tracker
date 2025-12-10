package task

import (
	"fmt"
	"log"
	"os"

	"github.com/Archiker-715/Task-Tracker/constants"
)

func AddTask(taskDescription, taskStatus string) (err error) {
	var file *os.File
	defer file.Close()

	if ok, size := fileExists(constants.TasksFileName); !ok {
		log.Printf("%q not found, will be created in current directory\n", constants.TasksFileName)
		if file, err = createFile(constants.TasksFileName); err != nil {
			return fmt.Errorf("add task error: %w", err)
		}
		log.Printf("file %q successfully created", constants.TasksFileName)
		fillTask(file, taskDescription, taskStatus, size)
	} else {
		fillTask(openFile(constants.TasksFileName), taskDescription, taskStatus, size)
	}

	return nil
}
