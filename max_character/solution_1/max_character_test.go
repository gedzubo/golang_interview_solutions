package main

import "testing"

func TestMaxCharacter(t *testing.T) {
	tables := []struct {
		input  string
		result string
	}{
		{"abbba", "b"},
		{"aabba", "a"},
		{"ababbaa", "a"},
		{"ababa", "a"},
	}

	for _, table := range tables {
		result := MaxCharacter(table.input)
		if result != table.result {
			t.Errorf(
				"MaxCharacter('%s') - expected: %s, got: %s",
				table.input,
				table.result,
				result,
			)
		}
	}
}
