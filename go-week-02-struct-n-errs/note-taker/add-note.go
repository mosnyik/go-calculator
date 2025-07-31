package notetaker

import (
	"fmt"
	// "os"
	// "bufio"
	// "strings"
)


func AddNote() {
	// what I need to display todos
	// 1. S/N #
	// 2. Category
	// 3. Task
	// 4. Done?
	// 5. CreatedAt
	// 6. CompletedAt

	// CHALLENGES
	// 1. I have issues reading a full link or space delimited strings using Println
	// solution:- I would use an array to save the questions
	// 			- I would pull the questions from the array
	// 			- I would loop through the questions and print them one after the other
	
	questions := [2]string{
		"Enter the task name: ",
		"Enter the task category: ",
	}
	var task string
	var category string

	for i, questions := range questions{
		fmt.Print(questions)
		if i == 1{
			fmt.Scan(&task)
		} else if i == 2 {
			fmt.Scan(&category)
		}
		
		
	}
		fmt.Println("Adding new task...", task, "into ", category, " category")

	// // scan := bufio.NewScanner(os.Stdin)
   
	// fmt.Println("Enter a new task:")
	//  task := os.Args[1:]
	// fmt.Scan(&task)

	// var category string

	// fmt.Println("Enter the task category:")
	// fmt.Scan(&category)




	

	fmt.Println("New note added successfully!")
}