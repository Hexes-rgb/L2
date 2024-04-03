package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSpace(input)

		if input == "\\quit" {
			break
		}

		commands := strings.Split(input, "|")
		if len(commands) > 1 {
			executePipes(commands)
		} else {
			executeCommand(input)
		}
	}
}

func executePipes(commands []string) {
	var previousStdout *io.ReadCloser
	for i, commandStr := range commands {
		commandStr = strings.TrimSpace(commandStr)
		parts := strings.Fields(commandStr)
		command := exec.Command(parts[0], parts[1:]...)

		var err error
		if previousStdout != nil {
			command.Stdin = *previousStdout
		}

		if i < len(commands)-1 {
			var stdout io.ReadCloser
			stdout, err = command.StdoutPipe()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Command", err)
				return
			}
			previousStdout = &stdout
		} else {
			command.Stdout = os.Stdout
		}

		err = command.Start()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error starting Command", err)
			return
		}

		if previousStdout != nil && i > 0 {
			err = (*previousStdout).Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error closing previous StdoutPipe", err)
				return
			}
		}
	}
}

func executeCommand(input string) {
	parts := strings.Fields(input)
	switch parts[0] {
	case "cd":
		if len(parts) < 2 {
			fmt.Println("Expected argument to 'cd'")
			return
		}
		err := os.Chdir(parts[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println(dir)
		}
	case "echo":
		fmt.Println(strings.Join(parts[1:], " "))
	case "kill":
		if len(parts) < 2 {
			fmt.Println("Expected argument to 'kill'")
			return
		}
		pid, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid process ID")
			return
		}
		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		err = process.Kill()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	case "ps":
		cmd := exec.Command("ps", "aux")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	default:
		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
