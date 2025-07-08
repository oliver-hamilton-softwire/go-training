package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	name := getName()
	fmt.Println(getOutput(name))
}

func getOutput(name string) string {
	// Exercise 6
	if name == "" {
		return "Ok, no greeting for you"
	}
	return greeting(name)
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	// Check that the characters at opposite ends of the string match
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func recogniseCreators(name string, greetingStr string) string {
	// Exercise 1 & 3
	creators := []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}

	for _, creator := range creators {
		if name == creator {
			greetingStr += " Thanks for creating me!"
		}
	}
	return greetingStr
}

func greeting(name string) string {
	// Exercise 4
	firstName := strings.Split(name, " ")[0]
	greetingStr := "Hello, " + firstName + "."

	greetingStr = recogniseCreators(name, greetingStr)

	// Exercise 2
	if len(firstName) > 20 {
		greetingStr = "Hello, " + firstName[:20] + "... Wow, that name's too long for me!"
	}

	// Exercise 5
	if isPalindrome(name) {
		greetingStr += " Cool, a palindromic name!"
	}
	return greetingStr
}
