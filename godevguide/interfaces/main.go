package main

import "fmt"

// new type called bot
// if you are a type in this program with a function called 'getGreeting' that returns a string
// then you are now also of type bot, and you have access to functions of this type
type bot interface {
	getGreeting() string
}

// this printGreeting function is available to bot types and any other type that implements the bot interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	// Imagine very custom logic for generating a greeting here
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
