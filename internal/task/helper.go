package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Archiker-715/Task-Tracker/constants"
	fm "github.com/Archiker-715/Task-Tracker/internal/file-manager"
)

var (
	ErrFileNotExists = fmt.Errorf("%q not found, please add your first task or create file", constants.TasksFileName)
	ErrEmptyFile     = fmt.Errorf("%q is empty, please add your first task", constants.TasksFileName)
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

	if ok, size := fm.FileExists(constants.TasksFileName); !ok {
		return ErrFileNotExists
	} else if ok && size == 0 {
		return ErrEmptyFile
	}

	return nil
}

func CheckErr(err error) error {
	if errors.Is(err, ErrFileNotExists) {
		return fmt.Errorf("list tasks err: %w", err)
	} else if errors.Is(err, ErrEmptyFile) {
		return fmt.Errorf("list tasks err: %w", err)
	}
	return nil
}

func Unquote(v string) (string, error) {
	var unquotedStr string
	if strings.Contains(v, `"`) {
		firstIdx := strings.Index(v, `"`)
		if firstIdx == -1 {
			return "", fmt.Errorf("the task's description could not be recognized")
		}
		firstPartString := v[firstIdx+1:]

		secondIdx := strings.Index(firstPartString, `"`)
		if secondIdx == -1 {
			return "", fmt.Errorf("the task's description could not be recognized")
		}
		unquotedStr = firstPartString[:secondIdx]
	}
	return unquotedStr, nil
}
