package kata

func Disemvowel(comment string) string {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	s := ""

	for _, char := range comment {
		if !vowels[char] {
			s += string(char)
		}
	}

	return s
}
