package main

import "fmt"

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func makePalindrome(s string) {
	chars := []rune(s)
	l := len(chars)
	count := 0
	for i := 0; i < l/2; i++ {
		count += abs(int(chars[i]) - int(chars[l-i-1]))
	}
	fmt.Println(count)
}

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	for i := 0; i < n; i++ {

		var s string
		fmt.Scanf("%v\n", &s)

		makePalindrome(s)
	}
}
