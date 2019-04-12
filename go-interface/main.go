package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	//very custom english logic
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	//very custom spanish logic
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
