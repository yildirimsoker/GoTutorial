package main

import (
	"fmt"
	"log"

	"github.com/yildirimsoker/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")
	var myVar helpers.SomeType
	myVar.TypeName = "T Name"
	myVar.TypeNumber = 3

	fmt.Println(myVar)
}
