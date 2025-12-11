package task

import (
	"encoding/json"
	"log"
	"os"

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
