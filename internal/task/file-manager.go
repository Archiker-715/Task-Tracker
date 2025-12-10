package task

import (
	"fmt"
	"io"
	"log"
	"os"
)

func fileExists(fileName string) (bool, int) {
	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false, 0
	} else if err != nil {
		log.Fatalf("check file exists error: %v", err)
	}
	return true, int(fileInfo.Size())
}

func createFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("create file error: %w", err)
	}
	return file, nil
}

func openFile(fileName string) *os.File {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	return file
}

func readFile(file *os.File) []byte {
	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	return fileContent
}

func seekPosition(file *os.File) {
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("seek position error: %v", err)
	}
}
