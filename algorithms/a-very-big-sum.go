package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%v\n", &size)
	var sum int64
	for i := 0; i < size; i++ {
		var elem int64
		fmt.Scanf("%v", &elem)
		sum += elem
	}
	fmt.Println(sum)
}
