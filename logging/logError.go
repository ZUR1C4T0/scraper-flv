package logging

import (
	"log"
	"os"
	"path/filepath"
)

func LogError(path string, err error) {
	absPath, _ := filepath.Abs(path)
	file, _ := os.OpenFile(absPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	log.SetOutput(file)
	log.Println(err)
}
