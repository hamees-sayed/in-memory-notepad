package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	notes := []string{"Notes 1", "Notes 2", "Notes 3", "Notes 4"}
	var noteLength int

	fmt.Print("Enter the maximum number of notes:")
	fmt.Scan(&noteLength)

	for {
		fmt.Print("Enter a command and data:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		line := scanner.Text()
		arr := strings.SplitN(line, " ", 2)

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
	if len(notes) == 0 {
		return
	}

	for index, value := range notes {
		if value == "" {
			continue
		}

		fmt.Printf("[Info] %d: %v\n", index+1, value)
	}
}

func addNote(note string, notes *[]string, noteLength int) {
	if len(*notes) >= noteLength {
		fmt.Println("[Error] Notepad is full")
	} else {
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
