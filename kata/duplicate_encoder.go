package kata

import (
	"strings"
)

func DuplicateEncode(word string) string {
	word = strings.ToLower(word)
	occurences := make(map[rune]int)

	for _, char := range word {
		_, present := occurences[char]
		if !present {
			occurences[char] = 0
		}

		occurences[char]++
	}

	result := ""

	for _, char := range word {
		count := occurences[char]

		if count == 1 {
			result += "("
		} else {
			result += ")"
		}
	}

	return result
}
