package goIdioms

import (
	"errors"
	"log"
)


func ErrorExamples(){

	user, err := getUser(0)
	if err != nil {
		log.Printf("Main error: %v", err.Error())
	}
	err = conbineDoSomething(user)
	if err != nil {
		return
	}

	
}

func conbineDoSomething(user string) error {
	err := doSomething(user)
	if err != nil{
		return err
	}

	// do something again
	err = doSomething(user)
	if err != nil{
		return err
	}
return nil
}

func doSomething(user string) error{
	return nil
}



func getUser(id int) (string, error){
	if id <= 0{
		return "", errors.New("invalid user ID")
	}

	return  "The Username", nil
}