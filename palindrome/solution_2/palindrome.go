package main

// Palindrome("abba") => true
// Palindrome("abbba") => false
func Palindrome(input string) bool {
	inputIsPalindrome := true
	length := len(input)

	for i := 0; i < length; i++ {
		if input[i] != input[length-i-1] {
			inputIsPalindrome = false
			break
		}
	}

	return inputIsPalindrome
}
