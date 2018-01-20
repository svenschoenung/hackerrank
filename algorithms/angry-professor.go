package main

import "fmt"

func test() {
	var students, threshold int
	fmt.Scanf("%d %d\n", &students, &threshold)

	onTime := 0
	for i := 0; i < students; i++ {
		var arrival int
		fmt.Scanf("%d", &arrival)
		if arrival <= 0 {
			onTime++
		}
	}
	if onTime < threshold {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var tests int
	fmt.Scanf("%v\n", &tests)

	for i := 0; i < tests; i++ {
		test()
	}
}
