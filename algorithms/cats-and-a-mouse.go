package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func main() {
	var queries int
	fmt.Scanf("%d\n", &queries)

	for i := 0; i < queries; i++ {
		var x, y, z int
		fmt.Scanf("%d %d %d\n", &x, &y, &z)

		var distX = abs(x - z)
		var distY = abs(y - z)

		if distX < distY {
			fmt.Printf("Cat A\n")
		} else if distX > distY {
			fmt.Printf("Cat B\n")
		} else {
			fmt.Printf("Mouse C\n")
		}
	}
}
