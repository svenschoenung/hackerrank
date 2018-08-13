package main

import "fmt"

func split(s string) ([]rune, []rune) {
	chars := []rune(s)
	return chars[:len(chars)/2], chars[len(chars)/2:]
}

func charCount(chars []rune) map[rune]int {
	count := map[rune]int{}
	for _, c := range chars {
		count[c]++
	}
	return count
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func main() {
	var s1 string
	fmt.Scanf("%v\n", &s1)

	var s2 string
	fmt.Scanf("%v\n", &s2)

	chars1 := []rune(s1)
	chars2 := []rune(s2)

	count1 := charCount(chars1)
	count2 := charCount(chars2)

	diff := 0
	diffed := map[rune]bool{}
	for c, _ := range count1 {
		if !diffed[c] {
			diff += abs(count1[c] - count2[c])
			diffed[c] = true
		}
	}
	for c, _ := range count2 {
		if !diffed[c] {
			diff += abs(count1[c] - count2[c])
			diffed[c] = true
		}
	}

	fmt.Println(diff)
}
