package task

import (
	"fmt"
	"log"
	"os"

	"github.com/Archiker-715/Task-Tracker/constants"
	fm "github.com/Archiker-715/Task-Tracker/internal/file-manager"
)

type Tasks struct {
	Tasks []Task `json:"Tasks"`
}

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (t *Tasks) AddTask(taskDescription string) (err error) {
	var file *os.File
	defer file.Close()

	if ok, size := fm.FileExists(constants.TasksFileName); !ok {
		log.Printf("%q not found, will be created in current directory\n", constants.TasksFileName)
		if file, err = fm.CreateFile(constants.TasksFileName); err != nil {
			return fmt.Errorf("add task error: %w", err)
		}
		log.Printf("file %q successfully created", constants.TasksFileName)
		t.addTask(file, taskDescription, size)
	} else {
		t.addTask(fm.OpenFile(constants.TasksFileName), taskDescription, size)
	}

	return nil
}

func (t *Tasks) UpdateTask(taskId int, data string) error {
	if err := checkFileExist(constants.TasksFileName); err != nil {
		return fmt.Errorf("check file exist err: %w", err)
	}

	if err := t.updateTask(fm.OpenFile(constants.TasksFileName), taskId, data); err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	return nil
}

func (t *Tasks) DeleteTask(taskId int) error {
	if err := checkFileExist(constants.TasksFileName); err != nil {
		return fmt.Errorf("check file exist err: %w", err)
	}

	if err := t.deleteTask(fm.OpenFile(constants.TasksFileName), taskId); err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return nil
}
