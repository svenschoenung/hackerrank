package main

import "fmt"
import "math"

func main() {
	var size int
	fmt.Scanf("%v\n", &size)

	var max int64 = math.MinInt64
	var count int = 0

	for i := 0; i < size; i++ {
		var val int64
		fmt.Scanf("%v", &val)

		if val > max {
			max = val
			count = 1
		} else if val == max {
			count++
		}
	}

	fmt.Println(count)
}
