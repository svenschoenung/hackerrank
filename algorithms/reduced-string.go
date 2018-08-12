package main

import "fmt"
import "strings"

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

func main() {
	var s string
	fmt.Scanf("%v\n", &s)

	chars := charsInString(s)

	for true {
		reduced := s
		for _, c := range chars {
			c2 := string(c) + string(c)
			reduced = strings.Replace(reduced, c2, "", -1)
		}
		if reduced == s {
			break
		}
		s = reduced
	}
	if s == "" {
		fmt.Println("Empty String")
	} else {
		fmt.Println(s)
	}
}
