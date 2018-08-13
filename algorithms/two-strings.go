package main

import "fmt"
import "strings"

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	for i := 0; i < n; i++ {
		var s1 string
		fmt.Scanf("%v\n", &s1)

		var s2 string
		fmt.Scanf("%v\n", &s2)

		result := "NO"
		for c := 'a'; c <= 'z'; c++ {
			if strings.ContainsRune(s1, c) && strings.ContainsRune(s2, c) {
				result = "YES"
				break
			}
		}
		fmt.Println(result)
	}
}
