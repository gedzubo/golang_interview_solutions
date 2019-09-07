package main

// Palindrome("abba") => true
// Palindrome("abbba") => false
func Palindrome(input string) bool {
	reversed := ""

	for _, r := range input {
		character := string(r)
		reversed = character + reversed
	}

	return reversed == input
}
