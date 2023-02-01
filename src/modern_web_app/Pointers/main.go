package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println(myString)
	changeUsingPointers(&myString)
	log.Println(myString)
}

func changeUsingPointers(s *string) {
	log.Println("s is set to", s)
	newString := "Green"
	*s = newString
}
