package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%v\n", &size)

	positiveCount := 0
	negativeCount := 0
	zeroCount := 0

	for i := 0; i < size; i++ {
		var elem int
		fmt.Scanf("%v", &elem)
		if elem > 0 {
			positiveCount++
		} else if elem < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	fmt.Printf("%0.6f\n", float32(positiveCount)/float32(size))
	fmt.Printf("%0.6f\n", float32(negativeCount)/float32(size))
	fmt.Printf("%0.6f\n", float32(zeroCount)/float32(size))
}
