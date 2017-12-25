package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%v\n", &size)

	diagonalSum1 := 0
	diagonalSum2 := 0

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			var elem int
			fmt.Scanf("%v", &elem)
			if i == j {
				diagonalSum1 += elem
			}
			if i+j == size-1 {
				diagonalSum2 += elem
			}
		}
	}

	if diagonalSum1 > diagonalSum2 {
		fmt.Println(diagonalSum1 - diagonalSum2)
	} else {
		fmt.Println(diagonalSum2 - diagonalSum1)
	}
}
