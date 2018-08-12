package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%v\n", &s)

	count := 0
	for i, char := range s {
		if (i%3 == 0 || i%3 == 2) && char != 'S' {
			count++
		}
		if i%3 == 1 && char != 'O' {
			count++
		}
	}
	fmt.Println(count)
}
