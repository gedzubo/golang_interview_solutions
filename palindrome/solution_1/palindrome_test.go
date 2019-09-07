package main

import "testing"

func TestPalindrome(t *testing.T) {
	tables := []struct {
		input  string
		result bool
	}{
		{"abba", true},
		{"abbba", true},
		{"aabba", false},
		{"ababba", false},
		{"ababa", true},
	}

	for _, table := range tables {
		result := Palindrome(table.input)
		if result != table.result {
			t.Errorf(
				"Palindrome('%s') - expected: %v, got: %v",
				table.input,
				table.result,
				result,
			)
		}
	}
}
