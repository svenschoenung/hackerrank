package main

import "fmt"

func charsInString(s string) []rune {
	hasChar := map[rune]bool{}
	for _, char := range s {
		hasChar[char] = true
	}
	var chars []rune
	for char, _ := range hasChar {
		chars = append(chars, char)
	}
	return chars
}

func removeChars(s string, c1 rune, c2 rune) string {
	var str string
	for _, char := range s {
		if char == c1 || char == c2 {
			str += string(char)
		}
	}
	return str
}

func isValid(s string) bool {
	chars := []rune(s)
	for i, char := range chars {
		if i > 0 && char == chars[i-1] {
			return false
		}
	}
	return true
}

func main() {
	var slen int
	fmt.Scanf("%v\n", &slen)

	var s string
	fmt.Scanf("%v\n", &s)

	chars := charsInString(s)

	maxLen := 0
	for _, c1 := range chars {
		for _, c2 := range chars {
			if c1 != c2 {
				s2 := removeChars(s, c1, c2)
				l := len([]rune(s2))
				if l > maxLen && isValid(s2) {
					maxLen = l
				}
			}
		}
	}
	fmt.Println(maxLen)
}
