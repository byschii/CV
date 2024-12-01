package fsutils

import (
	"log"
	"os"
)

// createOrClearFile creates a new file or clears the existing file and returns *os.File.
func CreateOrClearFile(fileName string) *os.File {
	// Open the file with O_RDWR|O_CREATE|O_TRUNC flags to create or truncate
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Failed to open or create file: %v", err)
	}
	return file
}
