package main

func main() {
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

	//var colors map[string]string
	//colors := make(map[int]string)
	//colors[10] = "Red"
	//delete(colors, 10)
	//fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		println("color is " + color + " hex is " + hex)
	}
}
