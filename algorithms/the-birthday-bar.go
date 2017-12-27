package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d\n", &size)

	bar := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &bar[i])
	}

	var d, m int
	fmt.Scanf("\n%d", &d)
	fmt.Scanf("%d\n", &m)

	count := 0
	for i := 0; i < size - m + 1; i++ {
		sum := 0
		for j := i; j < i + m; j++ {
			sum += bar[j]
		}
		if sum == d {
			count++
		}
	}
	fmt.Println(count)
}
