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

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%v\n", &s)

		w1, w2 := split(s)
		if len(w1) != len(w2) {
			fmt.Println(-1)
		} else {
			count1 := charCount(w1)
			count2 := charCount(w2)

			diffs := map[rune]int{}
			for c, _ := range count1 {
				diffs[c] = count1[c] - count2[c]
			}
			for c, _ := range count2 {
				diffs[c] = count1[c] - count2[c]
			}
			diff := 0
			for _, count := range diffs {
				if count > 0 {
					diff += count
				}
			}
			fmt.Println(diff)
		}
	}
}
