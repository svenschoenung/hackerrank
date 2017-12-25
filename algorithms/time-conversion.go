package main

import "fmt"

func main() {
	var h, m, s int
	var format string
	fmt.Scanf("%d:%d:%d%s\n", &h, &m, &s, &format)

	if format == "PM" && h < 12 {
		fmt.Printf("%02d:%02d:%02d\n", h+12, m, s)
	} else if format == "AM" && h == 12 {
		fmt.Printf("%02d:%02d:%02d\n", h-12, m, s)
	} else {
		fmt.Printf("%02d:%02d:%02d\n", h, m, s)
	}
}
