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
	greetingStr := greeting(name)
	if name == "Robert Griesemer" {
		greetingStr += ". Thanks for creating me!"
	}
	fmt.Println(greetingStr)
}

func greeting(name string) string {
	return "Hello, " + name
}
