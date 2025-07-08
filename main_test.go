package main

import "testing"

type Test struct {
	name     string
	expected string
}

var greetingTests = []Test{
	{"Alice", "Hello, Alice."},
	{"Bob", "Hello, Bob."},
	{"Robert Griesemer", "Hello, Robert. Thanks for creating me!"},
	{"Ken Thompson", "Hello, Ken. Thanks for creating me!"},
	{"Rob Pike", "Hello, Rob. Thanks for creating me!"},
	{"aaaaaaaaaabbbbbbbbbba", "Hello, aaaaaaaaaabbbbbbbbbb... Wow, that name's too long for me!"},
	{"Muhammad Ali", "Hello, Muhammad."},
	{"aba", "Hello, aba. Cool, a palindromic name!"},
	{"ababababababababababa", "Hello, abababababababababab... Wow, that name's too long for me! Cool, a palindromic name!"},
}

var outputTests = []Test{
	{"", "Ok, no greeting for you"},
	{"Alice", "Hello, Alice."},
}

func TestGreeting(t *testing.T) {
	for _, test := range greetingTests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting, for name %s got: %s expected: %s", test.name, result, test.expected)
		}
	}
}

func TestGetOutput(t *testing.T) {
	for _, test := range outputTests {
		result := getOutput(test.name)
		if result != test.expected {
			t.Errorf("incorrect output, for name %s got: %s expected: %s", test.name, result, test.expected)
		}
	}
}
