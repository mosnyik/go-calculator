package notetaker

import (
	"fmt")


func AddNote() {
	fmt.Println("Enter a new note:")
	var task string
	fmt.Scan(&task)

	fmt.Println("Adding new note...", task)

	fmt.Println("New note added successfully!")
}