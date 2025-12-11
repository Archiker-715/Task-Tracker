package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Archiker-715/Task-Tracker/constants"
	fm "github.com/Archiker-715/Task-Tracker/internal/file-manager"
)

func (t Tasks) findMaxId() int {
	var maxId int

	if len(t.Tasks) > 0 {
		maxId = t.Tasks[0].Id
		for _, task := range t.Tasks {
			if task.Id > maxId {
				maxId = task.Id
			}
		}
		return maxId
	}
	return 0
}

func (t Tasks) writeFile(file *os.File) {
	defer file.Close()
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatalf("marshalling err: %v", err)
	}

	fm.SeekStartPos(file)
	if err := file.Truncate(int64(fm.SeekCurrentPos(file))); err != nil {
		log.Fatalf("failed to truncate file: %v", err)
	}

	if _, err := file.Write(b); err != nil {
		log.Fatalf("writing task error: %v", err)
	}
}

func unmarshallFile(file *os.File, v interface{}) {
	if err := json.Unmarshal(fm.ReadFile(file), v); err != nil {
		log.Fatalf("unmarshalling err: %v", err)
	}
}

func checkFileExist(fileName string) error {
	var (
		fileNotExists = fmt.Errorf("%q not found, please add your first task", constants.TasksFileName)
		emptyFile     = fmt.Errorf("%q is empty, please add your first task", constants.TasksFileName)
	)

	if ok, size := fm.FileExists(constants.TasksFileName); !ok {
		return fileNotExists
	} else if ok && size == 0 {
		return emptyFile
	}

	return nil
}
