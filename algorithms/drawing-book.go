package main

import "fmt"

func main() {
	var pages int
	fmt.Scanf("%d\n", &pages)

	var page int
	fmt.Scanf("%d\n", &page)

	fromFront := (page - (page % 2)) / 2
	fromBack := (pages-(pages%2))/2 - fromFront

	if fromFront < fromBack {
		fmt.Printf("%d\n", fromFront)
	} else {
		fmt.Printf("%d\n", fromBack)
	}
}
