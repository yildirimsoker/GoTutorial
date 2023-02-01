package main

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}
type turkishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	tb := turkishBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(tb)
}

func printGreeting(b bot) {
	println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Halo"
}
func (turkishBot) getGreeting() string {
	return "Selam"
}
