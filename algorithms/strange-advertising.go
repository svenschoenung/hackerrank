package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	liked := 2
	shared := 5
        for i := 1; i < n; i++ {
		shared = shared / 2 * 3
		liked += shared / 2
	}
	fmt.Printf("%d\n", liked )
}
