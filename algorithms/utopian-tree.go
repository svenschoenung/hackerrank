package main

import "fmt"

func height(cycles int) int {
	height := 1
	for i := 1; i <= cycles; i++ {
		if i%2 == 1 {
			height = height * 2
		} else {
			height++
		}
	}
	return height
}

func main() {
	var tests int
	fmt.Scanf("%v\n", &tests)

	for i := 0; i < tests; i++ {
		var cycles int
		fmt.Scanf("%d", &cycles)
		fmt.Printf("%d\n", height(cycles))
	}
}
