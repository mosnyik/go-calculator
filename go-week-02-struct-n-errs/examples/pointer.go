package goIdioms

import (
	"fmt"
)

	func modifyName(n *string) {
		*n = "Mosnyik"
	}

	func (u *User) modifyUserName(newName string){
		u.name = newName
	}

	func modUserName(u User, name string){
		u.name = name
		fmt.Println("Modified user in modUserName()", u)
	}
	type User struct{
		id int
		name string
	}

func PointerExamples(){
	name := "Moses"

	fmt.Println("Name before modification:", name)

	modifyName(&name)
	fmt.Println("Name after modification:", name)

	// define a variable userID
	userID := 123
	fmt.Println("userID before modifying", userID)
	// reference the userID variable using a pointer
	anotherUserID := &userID
	// modify the variable using a pointer
	*anotherUserID=456

	fmt.Println("userID", userID)
	fmt.Println("userID address", &userID)
	fmt.Println("anotherUserID address", &anotherUserID)
	fmt.Println("But anotherUserID", anotherUserID)

	// defining a type point
	var id *int

	fmt.Println("The ID", id)

	id = &userID
    fmt.Println("The ID after change", *id)

	age := new(int)
	println("The age", age)
	*age = 30
	fmt.Println("The age after change", *age)


	user := User{
		id: 1,
		name: "Moses",
	}
	user.modifyUserName("Mosnyik");
	modUserName(user, "Mosnyik")
	fmt.Println("User name modified:", user.name)
} 