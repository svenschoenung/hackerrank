package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	max := 0
	counts := make(map[int]int)
	for i := 0; i < n; i++ {
		var val int
		fmt.Scanf("%d", &val)
		counts[val]++
		if counts[val] > max {
			max = counts[val]
		}
	}

	fmt.Printf("%d\n", n-max)
}
