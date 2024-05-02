//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	CmdHello = "hello"
	CmdBye   = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numCommands := 0
	numLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		text := strings.TrimSpace(line)
		// if line == "Q" || line == "q" {
		if strings.ToLower(text) == "q" {
			break
		} else {
			switch text {
			case CmdHello:
				numCommands++
				fmt.Println("Command resposne: Hello")

			case CmdBye:
				fmt.Println("Command resposne: ", CmdBye)
				numCommands++
			}
			if text != "" {
				numLines++
			}
		}
	}
	fmt.Printf("Number of non-blank lines entered: %v\n", numLines)
	fmt.Printf("Number of commands entered: %v\n", numCommands)

}
