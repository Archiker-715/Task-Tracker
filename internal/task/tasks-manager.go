package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Archiker-715/Task-Tracker/constants"
	fm "github.com/Archiker-715/Task-Tracker/internal/file-manager"
)

func (t *Tasks) addTask(file *os.File, taskDescription string, fileSize int) {

	if fileSize > 0 {
		unmarshallFile(file, t)
	}

	t.Tasks = append(t.Tasks, Task{
		Id:          (*t).findMaxId() + 1,
		Description: taskDescription,
		Status:      constants.Todo,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	})

	t.writeFile(file)
}

func (t *Tasks) listTasks(file *os.File) error {

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

func (t *Tasks) filteredListTasks(file *os.File, filter string) error {

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

	// var jsonData interface{}
	// err := json.Unmarshal(fm.ReadFile(file), &jsonData)
	// if err != nil {
	// 	return fmt.Errorf("err while parsing file: %w", err)
	// }

	formattedJSON, err := json.MarshalIndent(outputTasks, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling indent err: %w", err)
	}

	fmt.Println(string(formattedJSON))

	return nil
}

func (t *Tasks) updateTask(file *os.File, taskId int, data string) error {
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

func (t *Tasks) deleteTask(file *os.File, taskId int) error {
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
