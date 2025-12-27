package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		cwd, err := os.Getwd()
		if err != nil {
			cwd = "unknown"
		}

		fmt.Printf("\033[32m%s\033[0m > ", cwd)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execCommand(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execCommand(input string) error {
	input = strings.Trim(input, "\n")
	args := strings.Split(input, " ")

	leadingArg := args[0]
	switch leadingArg {
	case "cd":
		if len(args) < 2 {
			return errors.New("please add directory after `cd`")
		}

		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	default:
		cmd := exec.Command(args[0], args[1:]...)

		// Set the correct output device.
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		// Execute the command and return the error.
		return cmd.Run()
	}

	return nil
}
