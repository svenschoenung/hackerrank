package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	clouds := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &clouds[i])
	}

	e := 100
	i := 0
	jumps := 0
	for i != 0 || jumps == 0 {
		i = (i + k) % n
		e -= 1
		if clouds[i] == 1 {
			e -= 2
		}
		jumps++
	}
	fmt.Printf("%d\n", e)
}
