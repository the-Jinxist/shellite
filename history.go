package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

var _ CommandHistory = &FileCommandHistory{}

type CommandHistory interface {
	Up(prefix string) string
	Down(prefix string) string
	SaveCommandToHistroy(command string) error
}

type FileCommandHistory struct {
	history []string
	file    *os.File
	index   int
}

func NewCommandHistory(file *os.File) *FileCommandHistory {

	var history []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		history = append(history, line)
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Fatalf("error while scanning from file")
	}

	// Parse line by line separated by \n

	//Save commands to string array
	return &FileCommandHistory{
		history: history,
		file:    file,
		index:   -1,
	}
}

// Down implements CommandHistory.
func (f *FileCommandHistory) Down(prefix string) string {
	if prefix == "" && f.index < len(f.history)-1 {
		f.index++
		return f.history[f.index]
	}

	return ""
}

// SaveCommandToHistroy implements CommandHistory.
func (f *FileCommandHistory) SaveCommandToHistroy(command string) error {
	f.history = append([]string{command}, f.history...)

	var value string
	for _, comm := range f.history {
		value += comm + "\n"
	}

	_, err := f.file.WriteString(value)
	return err
}

// Up implements CommandHistory.
func (f *FileCommandHistory) Up(prefix string) string {
	if prefix == "" {
		f.index--
		return f.history[f.index]
	}

	return ""
}
