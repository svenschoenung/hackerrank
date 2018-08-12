package main

import "fmt"

func isPalindrome(chars []rune) bool {
	l := len(chars)
	for i := 0; i < l/2; i++ {
		if chars[i] != chars[l-i-1] {
			return false
		}
	}
	return true
}

func makePalindrome(s string) int {
	chars := []rune(s)
	l := len(chars)
	for i := 0; i < l/2; i++ {
		if chars[i] != chars[l-i-1] {
			if isPalindrome(chars[i+1 : l-i]) {
				return i
			}
			if isPalindrome(chars[i : l-i-1]) {
				return l - i - 1
			}
			return -1
		}
	}
	return -1
}

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	for i := 0; i < n; i++ {

		var s string
		fmt.Scanf("%v\n", &s)

		index := makePalindrome(s)
		fmt.Println(index)
	}
}
