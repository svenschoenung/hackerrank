package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	maxType := -1
	maxCount := 0
	typeCounts := make(map[int]int)
	for i := 0; i < n; i++ {
		var birdType int
		fmt.Scanf("%d", &birdType)
		typeCounts[birdType]++
		if typeCounts[birdType] > maxCount {
			maxCount = typeCounts[birdType]
			maxType = birdType
		} else if typeCounts[birdType] == maxCount && birdType < maxType {
			maxType = birdType
		}
	}

	fmt.Println(maxType)
}
