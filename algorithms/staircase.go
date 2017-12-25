package main

import "fmt"
import "strings"

func main() {
	var size int
	fmt.Scanf("%v\n", &size)

	for i := 0; i < size; i++ {
		fmt.Print(strings.Repeat(" ", size-i-1))
		fmt.Println(strings.Repeat("#", i+1))
	}
}
