package main

// MaxCharacter("aaba") => "a"
// MaxCharacter("abbba") => "b"
func MaxCharacter(input string) string {
	characterMap := make(map[string]int)
	maxCharacter := ""
	maxValue := 0

	for _, r := range input {
		character := string(r)

		_, value := characterMap[character]
		if value {
			characterMap[character]++
		} else {
			characterMap[character] = 1
		}
	}

	for key, value := range characterMap {
		if value > maxValue {
			maxCharacter = key
			maxValue = value
		}
	}

	return maxCharacter
}
