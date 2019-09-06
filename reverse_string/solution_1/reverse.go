package main

func Reverse(original string) string {
	reversed := ""

	for _, r := range original {
		character := string(r)
		reversed = character + reversed
	}

	return reversed
}
