package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	jim := person{
		firstName: "jim",
		lastName:  "Alex",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9000,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateName("jimmy")
	jim.updateName("jimmy")
	jim.print()

	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)

	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"
	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	//deneme := 10
	//changeIntTest(&deneme)
	//println(deneme)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//func changeIntTest(value *int) {
//	(*value) = 21
//}
