package main

import (
	"bufio"
	"fmt"
	"os"
)

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	name := getName()
	fmt.Println(greeting(name))
}

func greeting(name string) string {
	greetingStr := "Hello, " + name[:20]
	// Exercise 1
	if name == "Robert Griesemer" {
		greetingStr += ". Thanks for creating me!"
	}
	// Exercise 2
	if len(name) > 20 {
		greetingStr += "... Wow, that name's too long for me!"
	}
	return greetingStr
}
