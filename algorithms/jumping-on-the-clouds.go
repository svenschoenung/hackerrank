package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	clouds := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &clouds[i])
	}

	jumps := 0
	i := 0
	for ; i < n - 1; {
		if i + 2 < n && clouds[i + 2] == 0 {
			i += 2
			jumps++
		} else if i + 1 < n && clouds[i + 1] == 0 {
			i += 1
			jumps++
		}
	}
	fmt.Printf("%d\n", jumps)
}
