package main

import "log"

type mySturct struct {
	FirstName string
}

func main() {
	var myVar mySturct
	myVar.FirstName = "Ali"

	myVar2 := mySturct{
		FirstName: "Ahmet",
	}

	log.Println("myVar is: ", myVar.getFirstName())
	log.Println("myVar2 is: ", myVar2.getFirstName())
}

func (m *mySturct) getFirstName() string {
	return m.FirstName
}
