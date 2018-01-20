package main

import "fmt"

func main() {
	var n, maxHeight int
	fmt.Scanf("%d %d\n", &n, &maxHeight)

	beverages := 0
	for i := 0; i < n; i++ {
		var height int
		fmt.Scanf("%d", &height)
		if height > maxHeight {
			beverages += height - maxHeight
			maxHeight = height
		}
	}

	fmt.Printf("%d\n", beverages)
}
