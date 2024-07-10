package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate time.Time
}

func (user *User) printFirstName() string {
	return user.FirstName
}

func main() {
	user := User{
		FirstName: "John",
		LastName:  "Oliver",
		BirthDate: time.Date(1994, 6, 12, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println("User: ", user)
	fmt.Println("User's first name: ", user.FirstName)
	fmt.Println("User's last name: ", user.LastName)
	fmt.Println("User's birth date: ", user.BirthDate.Format("02/01/2006"))
	fmt.Println("User's first name using method: ", user.printFirstName())

	// Format the birth, needs to take in consideration the time: 02/01/2006 15:04:05, even I don't have idea why this date
}

/*
	In Go, Captal letter is used to export a variable, function, etc. to other packages.
	Lower case is used to keep it private.
*/
