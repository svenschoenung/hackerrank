package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	inv_p := make([]int, n+1)
	for x := 1; x <= n; x++ {
		var px int
		fmt.Scanf("%d", &px)
		inv_p[px] = x
	}

	for x := 1; x <= n; x++ {
		fmt.Printf("%d\n", inv_p[inv_p[x]])
	}
}
