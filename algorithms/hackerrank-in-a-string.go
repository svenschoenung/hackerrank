package main

import "fmt"

func test(s string) {
	chars := []rune("hackerrank")
	i := 0
	for _, char := range s {
		if char == chars[i] {
			i++
			if i >= 10 {
				fmt.Println("YES")
				return
			}
		}
	}
	fmt.Println("NO")
}

func main() {
	var q int
	fmt.Scanf("%v\n", &q)

	for i := 0; i < q; i++ {
		var s string
		fmt.Scanf("%v\n", &s)
		test(s)
	}
}
