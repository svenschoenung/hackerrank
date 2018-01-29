package main

import "fmt"


func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	min := n
	result := -1
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] == arr[j] {
				if j - i < min {
					min = j - i
					result = min
				}
			}
		}
	}
	fmt.Printf("%d\n", result)
}
