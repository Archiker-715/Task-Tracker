package task

import (
	"encoding/json"
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
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func (t *Tasks) AddTask(taskDescription string) (int, error) {

	var (
		file *os.File
		err  error
	)

	ok, _ := fm.FileExists(constants.TasksFileName)
	if !ok {
		log.Printf("%q not found, will be created in current directory\n", constants.TasksFileName)
		if file, err = fm.CreateFile(constants.TasksFileName); err != nil {
			return 0, fmt.Errorf("add taskfile error: %w", err)
		}
		t.writeFile(file)
		log.Printf("file %q successfully created", constants.TasksFileName)
	} else {
		file = fm.OpenFile(constants.TasksFileName)
	}
	defer file.Close()

	unmarshallFile(file, t)

	newTaskId := (*t).findMaxId() + 1

	t.Tasks = append(t.Tasks, Task{
		Id:          newTaskId,
		Description: taskDescription,
		Status:      constants.Todo,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	})

	t.writeFile(file)

	return newTaskId, nil
}

func (t *Tasks) ListTasks() error {

	if err := checkFileExist(constants.TasksFileName); err != nil {
		return fmt.Errorf("check file exist err: %w", err)
	}

	file := fm.OpenFile(constants.TasksFileName)
	defer file.Close()

	var jsonData interface{}
	err := json.Unmarshal(fm.ReadFile(file), &jsonData)
	if err != nil {
		return fmt.Errorf("err while parsing file: %w", err)
	}

	formattedJSON, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling indent err: %w", err)
	}

	fmt.Println(string(formattedJSON))

	return nil
}

func (t *Tasks) FilteredListTasks(filter string) error {

	if err := checkFileExist(constants.TasksFileName); err != nil {
		return fmt.Errorf("check file exist err: %w", err)
	}

	file := fm.OpenFile(constants.TasksFileName)
	defer file.Close()

	unmarshallFile(file, t)

	var outputTasks Tasks

	for i := range t.Tasks {
		if t.Tasks[i].Status == filter {
			outputTasks.Tasks = append(outputTasks.Tasks, t.Tasks[i])
		}
	}

	if len(outputTasks.Tasks) == 0 {
		fmt.Printf("Tasks with status %q not found\n", filter)
		return nil
	}

	formattedJSON, err := json.MarshalIndent(outputTasks, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling indent err: %w", err)
	}

	fmt.Println(string(formattedJSON))

	return nil
}

func (t *Tasks) UpdateTask(taskId int, data string) error {

	if err := checkFileExist(constants.TasksFileName); err != nil {
		return fmt.Errorf("check file exist err: %w", err)
	}

	file := fm.OpenFile(constants.TasksFileName)
	defer file.Close()

	unmarshallFile(file, t)

	err := func() error {
		for i := range t.Tasks {
			if t.Tasks[i].Id == taskId {
				var updated = time.Now().Format("2006-01-02 15:04:05")
				switch data {
				case constants.MarkInProgress:
					t.Tasks[i].Status = constants.InProgress
					t.Tasks[i].UpdatedAt = updated
					return nil
				case constants.MarkDone:
					t.Tasks[i].Status = constants.Done
					t.Tasks[i].UpdatedAt = updated
					return nil
				default:
					t.Tasks[i].Description = data
					t.Tasks[i].UpdatedAt = updated
					return nil
				}
			}
		}
		return fmt.Errorf("taskId %q not found", taskId)
	}()
	if err != nil {
		return fmt.Errorf("update task error: %w", err)
	}

	t.writeFile(file)

	return nil
}

func (t *Tasks) DeleteTask(taskId int) error {
	if err := checkFileExist(constants.TasksFileName); err != nil {
		return fmt.Errorf("check file exist err: %w", err)
	}

	file := fm.OpenFile(constants.TasksFileName)
	defer file.Close()
	unmarshallFile(file, t)

	err := func() error {
		for i := range t.Tasks {
			if t.Tasks[i].Id == taskId {
				leftTasks := t.Tasks[:i]
				rightTasks := t.Tasks[i+1:]
				outputTasks := append(leftTasks, rightTasks...)
				t.Tasks = outputTasks
				return nil
			}
		}
		return fmt.Errorf("taskId %q not found", taskId)
	}()
	if err != nil {
		return fmt.Errorf("update task error: %w", err)
	}

	t.writeFile(file)

	return nil
}
