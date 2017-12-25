package main

import "fmt"
import "math"

func main() {
	var arr [5]int64
	for i := 0; i < 5; i++ {
		fmt.Scanf("%v", &arr[i])
	}

	var min int64 = math.MaxInt64
	var max int64 = math.MinInt64

	for skip := 0; skip < 5; skip++ {
		var sum int64
		for i := 0; i < 5; i++ {
			if i != skip {
				sum += arr[i]
			}
		}
		if sum < min {
			min = sum
		}
		if sum > max {
			max = sum
		}
	}

	fmt.Printf("%d %d\n", min, max)
}
