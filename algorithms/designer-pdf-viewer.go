package main

import "fmt"

func main() {
	var heights [26]int
	for i := 0; i < 26; i++ {
		fmt.Scanf("%d", &heights[i])
	}

	var word string
	fmt.Scanf("%s", &word)

	width := len(word)
	maxHeight := 0
	for i := 0; i < len(word); i++ {
		letter := word[i]
		height := heights[letter-97]
		if height > maxHeight {
			maxHeight = height
		}
	}

	fmt.Printf("%d\n", width*maxHeight)
}
