package main

import "fmt"

type englishBot struct{}
type frenchBot struct{}

func main() {
	//eb := englishBot{}
	fb := frenchBot{}

	//printGreeting(eb)
	printGreeting(fb)
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi There!"
}

func (frenchBot) getGreeting() string {
	// Very custom logic for generating a French greeting
	return "Bonjour!"
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

func printGreeting(fb frenchBot) {
	fmt.Println(fb.getGreeting())
}
