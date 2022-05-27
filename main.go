package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var noteLength int

	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scan(&noteLength)

	// init string slice with specified number of elements
	notes := make([]string, 0, noteLength)
	fmt.Println()

	// "while(true)" program loop
	for {
		fmt.Print("Enter a command and data: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		// separate user input into:
		// [0] : command
		// [1] : information (needed for create function)
		arr := strings.SplitN(line, " ", 2)

		if len(arr) != 2 {
			arr = append(arr, "")
		}

		if arr[0] == "exit" {
			exitProgram()
			break
		} else if arr[0] == "list" {
			readNoteList(notes)
		} else if arr[0] == "create" {
			addNote(arr[1], &notes, noteLength)
		} else if arr[0] == "clear" {
			deleteAllNotes(&notes)
		} else {
			unknownCommandError()
		}
	}
}

func exitProgram() {
	fmt.Print("[Info] Bye!\n")
}

func readNoteList(notes []string) {
	// if there are no notes, terminate function early
	if len(notes) == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}

	// iterate through notes, using index as output label
	for index, value := range notes {
		if value == "" {
			continue
		}

		fmt.Printf("[Info] %d: %v\n", index+1, value)
	}
}

func addNote(note string, notes *[]string, noteLength int) {
	// remove leading and trailing whitespace from note (ensures cleaned input and no "empty" notes)
	newNote := strings.TrimSpace(note)

	// Notepad will not accept notes past specified limit
	if len(*notes) >= noteLength {
		fmt.Println("[Error] Notepad is full")
	} else {
		// prevent empty note input
		if newNote == "" {
			fmt.Println("[Error] Missing note argument")
			return
		}
		*notes = append(*notes, note)
		fmt.Println("[OK] The note was successfully created")
	}
}

func deleteAllNotes(notes *[]string) {
	*notes = nil
	fmt.Println("[OK] All notes were successfully deleted")
}

func unknownCommandError() {
	fmt.Println("[Error] Unknown command")
}
