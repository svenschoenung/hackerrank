package main

import "fmt"

func countOnHouse(tree int, s int, t int, fruits int) int {
	count := 0
	for i := 0; i < fruits; i++ {
		var d int
		fmt.Scanf("%d", &d)
		if tree+d >= s && tree+d <= t {
			count++
		}
	}
	return count
}

func main() {
	var s, t int
	fmt.Scanf("%d %d\n", &s, &t)

	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)

	var m, n int
	fmt.Scanf("%d %d\n", &m, &n)

	fmt.Println(countOnHouse(a, s, t, m))
	fmt.Println(countOnHouse(b, s, t, n))
}
