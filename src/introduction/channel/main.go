package main

import (
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://stackoverflow.com/",
		"https://go.dev/",
		"https://www.amazon.com/",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c)
	// }

	for l := range c {
		go checkLink(l, c)
	}
}

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		println(link, "might be down")
// 		c <- "Might be down I think"
// 		return
// 	}

// 	println(link, "is up!")
// 	c <- "Yep its up"
// }

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		println(link, "might be down")
		c <- link
		return
	}

	println(link, "is up!")
	c <- link
}
