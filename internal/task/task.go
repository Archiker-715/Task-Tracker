package task

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Archiker-715/Task-Tracker/constants"
	fm "github.com/Archiker-715/Task-Tracker/internal/file-manager"
)

type Tasks struct {
	Tasks []Task `json:"Tasks"`
}

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
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
	var (
		fileNotExists = fmt.Errorf("%q not found, please add your first task", constants.TasksFileName)
		emptyFile     = fmt.Errorf("%q is empty, please add your first task", constants.TasksFileName)
	)

	if ok, size := fm.FileExists(constants.TasksFileName); !ok {
		return fileNotExists
	} else if ok && size == 0 {
		return emptyFile
	}

	if err := t.updateTask(fm.OpenFile(constants.TasksFileName), taskId, data); err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	return nil
}
