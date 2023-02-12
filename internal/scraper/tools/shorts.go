package tools

import (
	"log"
	"os"
)

func CreateDir(path string) {
	log.Printf("Creating Directory: %v", path)
	// Create dir if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			log.Printf("Error: %v", err)
		}
	}
}
