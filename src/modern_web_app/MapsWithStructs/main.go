package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Yildirim",
		LastName:  "Soker",
	}

	myMap["me"] = me

	log.Println("map name:", myMap["me"].FirstName, "map surname:", myMap["me"].LastName)
}
