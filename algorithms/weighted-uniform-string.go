package main

import "fmt"

func weight(r rune) int {
	return int(r) - 96
}

func weights(s string) map[int]bool {
	runes := []rune(s)
	w := map[int]bool{}

	count := 1
	for i, r := range s {
		w[weight(r)] = true
		if i > 0 && runes[i-1] == r {
			count++
			w[weight(r)*count] = true
		} else {
			count = 1
		}
	}

	return w
}

func main() {
	var s string
	fmt.Scanf("%v\n", &s)

	var n int
	fmt.Scanf("%v\n", &n)

	weights := weights(s)

	for i := 0; i < n; i++ {
		var q int
		fmt.Scanf("%v\n", &q)

		if weights[q] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
