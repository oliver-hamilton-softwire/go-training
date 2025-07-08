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
	greetingStr := "Hello, " + name
	if name == "Robert Griesemer" {
		greetingStr += ". Thanks for creating me!"
	}
	return greetingStr
}
