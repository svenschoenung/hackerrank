package main

import "fmt"
import "unicode"

func main() {
	var s string
	fmt.Scanf("%v\n", &s)
	words := 1
	for _, char := range s {
		if unicode.IsUpper(char) {
			words++
		}
	}
	fmt.Println(words)
}
