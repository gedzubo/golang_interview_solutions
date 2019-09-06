package main

import "testing"

func TestReverse(t *testing.T) {
	tables := []struct {
		input  string
		result string
	}{
		{"Hello", "olleH"},
		{"How are you?", "?uoy era woH"},
		{"123456789", "987654321"},
	}

	for _, table := range tables {
		result := Reverse(table.input)
		if result != table.result {
			t.Errorf(
				"Reverse('%s') - expected: %s, got: %s",
				table.input,
				table.result,
				result,
			)
		}
	}
}
