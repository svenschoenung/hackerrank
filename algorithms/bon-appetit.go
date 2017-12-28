package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d\n", &k)

	var sum int
	for i := 0; i < n; i++ {
		var item int
		fmt.Scanf("%d", &item)
		if i != k {
			sum += item
		}
	}

	var charged int
	fmt.Scanf("%d", &charged)

	var actual = sum / 2
	if charged == actual {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Printf("%d\n", charged-actual)
	}
}
