package main

import "fmt"

func charCount(chars []rune) map[rune]int {
	count := map[rune]int{}
	for _, c := range chars {
		count[c]++
	}
	return count
}

func main() {
	var s string
	fmt.Scanf("%v\n", &s)

	chars := []rune(s)
	counts := charCount(chars)

	if len(chars)%2 == 0 {
		for _, count := range counts {
			if count%2 != 0 {
				fmt.Println("NO")
				return
			}
		}
		fmt.Println("YES")
	} else {
		oddChars := false
		for _, count := range counts {
			if count%2 != 0 {
				if !oddChars {
					oddChars = true
				} else {
					fmt.Println("NO")
					return
				}
			}
		}
		fmt.Println("YES")
	}
}
