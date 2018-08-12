package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func equal(slice1, slice2 []int) bool {
	for i, v := range slice1 {
		if slice2[i] != v {
			return false
		}
	}
	return true
}

func diffs(s string) []int {
	runes := []rune(s)
	diffs := make([]int, len(runes)-1)
	for i, v := range runes {
		if i > 0 {
			diffs[i-1] = abs(int(v) - int(runes[i-1]))
		}
	}
	return diffs
}

func reverse(s string) string {
	reversed := ""
	for _, c := range s {
		reversed = string(c) + reversed
	}
	return reversed
}

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	for i := 0; i < n; i++ {
		var s1 string
		fmt.Scanf("%v\n", &s1)

		s2 := reverse(s1)

		d1 := diffs(s1)
		d2 := diffs(s2)

		if equal(d1, d2) {
			fmt.Println("Funny")
		} else {
			fmt.Println("Not Funny")
		}
	}
}
