package main

import "fmt"

func alternating(s string) {
	runes := []rune(s)
	count := 0
	for i, r := range runes {
		if i > 0 && r == runes[i-1] {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%v\n", &s)
		if len(s) == 0 {
			fmt.Println("")
		} else {
			alternating(s)
		}
	}
}
