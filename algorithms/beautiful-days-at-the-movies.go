package main

import "fmt"
import "strconv"

func reversed(x int) int {
	str := strconv.Itoa(x);
	chars := make([]rune, len(str))
	for i, char := range str {
		chars[len(str)-1-i] = char
	}
	r, _ := strconv.Atoi(string(chars))
	return r
}

func main() {
	var i, j, k int
	fmt.Scanf("%d %d %d\n", &i, &j, &k)

	count := 0
        for n := i; n <= j; n++ {
		if (n - reversed(n)) % k == 0 {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}
