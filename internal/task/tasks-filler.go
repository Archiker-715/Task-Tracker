package task

import (
	"fmt"
	"os"
	"time"

	"github.com/Archiker-715/Task-Tracker/constants"
)

func (t *Tasks) addTask(file *os.File, taskDescription string, fileSize int) {

	if fileSize > 0 {
		unmarshallFile(file, t)
	}

	t.Tasks = append(t.Tasks, Task{
		Id:          (*t).findMaxId() + 1,
		Description: taskDescription,
		Status:      constants.Todo,
		CreatedAt:   time.Now(),
	})

	t.writeFile(file)
}

func (t *Tasks) updateTask(file *os.File, taskId int, data string) error {
	unmarshallFile(file, t)

	if err := t.updateTaskData(taskId, data); err != nil {
		return fmt.Errorf("failed update task: %w", err)
	}

	t.writeFile(file)

	return nil
}
