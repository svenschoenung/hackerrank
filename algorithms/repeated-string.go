package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var str string
	fmt.Scanf("%s\n", &str)

	var n int
	fmt.Scanf("%d\n", &n)

	strlen := len(str)
	repeated := n / strlen

	count := 0
	for i := 0; i < min(n, strlen); i++ {
		if str[i] == 'a' {
			count++
		}
	}
	if repeated > 0 {
		count = count * repeated
		for i := 0; i < (n - (n / repeated * repeated)); i++ {
			if str[i] == 'a' {
				count++
			}
		}
	}

	fmt.Printf("%d\n", count)
}
