package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%v\n", &size)
	var sum = 0
	for i := 0; i < size; i++ {
		var elem int
		fmt.Scanf("%v", &elem)
		sum += elem
	}
	fmt.Println(sum)
}
