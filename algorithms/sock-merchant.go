package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var socks = make(map[int]int)
	for i := 0; i < n; i++ {
		var sockType int
		fmt.Scanf("%d", &sockType)
		socks[sockType]++
	}

	var sockPairs int
	for _, count := range socks {
		sockPairs += count / 2
	}

	fmt.Println(sockPairs)
}
