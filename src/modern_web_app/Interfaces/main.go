package main

import (
	"fmt"
)

type Animal interface {
	Say() string
	NumberOfLeg() int
}
type Dog struct {
	Name string
	Age  int
}

type Gorilla struct {
	Name  string
	Color string
}

func main() {

	newDog := Dog{
		Name: "Comar",
		Age:  12,
	}

	newGorilla := Gorilla{
		Name:  "Şahap",
		Color: "Red",
	}

	PrintAnimalInfo(&newDog)
	PrintAnimalInfo(&newGorilla)

}

func PrintAnimalInfo(a Animal) {
	fmt.Println("the animal says", a.Say(), "and this animal legs", a.NumberOfLeg())
}

func (a *Dog) Say() string {
	return "Haw Haw"
}

func (a *Dog) NumberOfLeg() int {
	return 4
}

func (a *Gorilla) Say() string {
	return "Grow"
}

func (a *Gorilla) NumberOfLeg() int {
	return 2
}
