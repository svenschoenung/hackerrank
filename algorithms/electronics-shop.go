package main

import "fmt"

func main() {
	var money, n, m int
	fmt.Scanf("%d %d %d\n", &money, &n, &m)

	var keyboards = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &keyboards[i])
	}

	var drives = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &drives[i])
	}

	maxCost := -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			cost := keyboards[i] + drives[j]
			if cost <= money && cost > maxCost {
				maxCost = cost
			}
		}
	}

	fmt.Printf("%d\n", maxCost)
}
