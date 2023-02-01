package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	for i := 0; i <= 10; i++ {
		println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow"}

	for x, animal := range animals {
		log.Println(x, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	colors := make(map[string]string)
	colors["yellow"] = "#dsdfsdf"
	colors["green"] = "#sadad"

	for i, color := range colors {
		log.Println(i, color)
	}

	var users []User
	users = append(users, User{"Yildirim", "Soker"})
	users = append(users, User{"Kaan", "Söker"})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName)
	}
}
