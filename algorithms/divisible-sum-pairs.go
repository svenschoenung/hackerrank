package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d\n", &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum := arr[i] + arr[j]
			if sum%k == 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
