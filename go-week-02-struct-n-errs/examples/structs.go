package goIdioms

import (
	"encoding/json"
	"fmt"
)

// updating struct using struct function
func (c *Customer) updateEmail(newEmail string) {
		c.email = newEmail
	}

	// Defining an struct to be nested in another struct
type Address struct {
	street string 
	city string 
}
// Struct function for Address struct
func (a Address ) printAddress(){
	fmt.Println("Address:", a)
}
// Defining a struct

	type Customer struct {
		name string 
		age int
		Address
		email string 
	}
	
func StructsExamples() {
	address := Address{
			street: "123 Amino Kano Crescent",
			city: "Abuja",
		}
	firstCustomer := Customer {
		name : "Moses",
		age : 25,
		Address : address,
		email : "email@gmail.com",
	 	}

		firstCustomer.updateEmail("newemail")
		firstCustomer.printAddress();
	fmt.Println("I am accessing the struct name, address and email:", firstCustomer.name, firstCustomer.email, firstCustomer.street )

	// struct tags
	type User struct {
		Name string `json:"name"`
		Age int `json:"age"`
	}

	user1 := User{
		Name : "Mose",
		Age: 50,
	}

	jsonData, _ := json.Marshal(user1)
	fmt.Println("Serialized User is", string(jsonData))
}