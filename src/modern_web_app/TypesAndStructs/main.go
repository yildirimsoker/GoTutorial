package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Yildirim",
		LastName:    "Soker",
		PhoneNumber: "05056092920",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber, "Birth Date:", user.BirthDate)
}
