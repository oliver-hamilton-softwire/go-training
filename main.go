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
	creators := []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
	// Exercise 1 & 3
	for _, creator := range creators {
		if name == creator {
			greetingStr += ". Thanks for creating me!"
		}
	}
	// Exercise 2
	if len(name) > 20 {
		greetingStr = "Hello, " + name[:20] + "... Wow, that name's too long for me!"
	}
	return greetingStr
}
