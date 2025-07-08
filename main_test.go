package main

import "testing"

var tests = []struct {
	name     string
	expected string
}{
	{"Alice", "Hello, Alice"},
	{"Bob", "Hello, Bob"},
	{"Robert Griesemer", "Hello, Robert. Thanks for creating me!"},
	{"Ken Thompson", "Hello, Ken. Thanks for creating me!"},
	{"Rob Pike", "Hello, Rob. Thanks for creating me!"},
	{"aaaaaaaaaabbbbbbbbbba", "Hello, aaaaaaaaaabbbbbbbbbb... Wow, that name's too long for me!"},
	{"Muhammad Ali", "Hello, Muhammad"},
}

func TestGreeting(t *testing.T) {
	for _, test := range tests {
		result := greeting(test.name)
		if result != test.expected {
			t.Errorf("incorrect greeting, for name %s got: %s expected: %s", test.name, result, test.expected)
		}
	}
}
