package main

var _ CommandHistory = &FileCommandHistory{}

type CommandHistory interface {
	Up(prefix string) string
	Down(prefix string) string
	SaveCommandToHistroy() error
}

type FileCommandHistory struct {
	history []string
}

func NewCommandHistory(filePath string) *FileCommandHistory {
	//Load file

	// Parse line by line separated by \n

	//Save commands to string array
	return &FileCommandHistory{}
}

// Down implements CommandHistory.
func (f *FileCommandHistory) Down(prefix string) string {
	panic("unimplemented")
}

// SaveCommandToHistroy implements CommandHistory.
func (f *FileCommandHistory) SaveCommandToHistroy() error {
	panic("unimplemented")
}

// Up implements CommandHistory.
func (f *FileCommandHistory) Up(prefix string) string {
	panic("unimplemented")
}
