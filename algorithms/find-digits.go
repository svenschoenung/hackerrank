package main

import "fmt"
import "strconv"

func findDigits(n int) int {
	count := 0
	str := strconv.Itoa(n)
	for _, char := range str {
		d, _ := strconv.Atoi(string([]rune{char}))
		if d != 0 && n%d == 0 {
			count++
		}
	}
	return count
}

func main() {
	var tests int
	fmt.Scanf("%d\n", &tests)

	for i := 0; i < tests; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		fmt.Printf("%d\n", findDigits(n))
	}
}
