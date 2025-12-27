package main

import (
	"os"
)

func OpenFile(path string) (*os.File, error) {
	//Load or create new file
	return os.Create(path)
}

func CloseFile(file *os.File) {
	if file == nil {
		return
	}

	file.Close()
}
