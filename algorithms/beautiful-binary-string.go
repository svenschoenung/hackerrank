package main

import "fmt"
import "strings"

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	var s string
	fmt.Scanf("%v\n", &s)

	if n < 3 {
		fmt.Println(0)
	} else if n == 3 {
		if s == "010" {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	} else {
		count := 0
		for strings.Contains(s, "01010") {
			s = strings.Replace(s, "01010", "01110", 1)
			count++
		}
		for strings.Contains(s, "010") {
			if strings.Contains(s, "0010") {
				s = strings.Replace(s, "0010", "0110", 1)
				count++
			}
			if strings.Contains(s, "1010") {
				s = strings.Replace(s, "1010", "1110", 1)
				count++
			}
			if strings.Contains(s, "0100") {
				s = strings.Replace(s, "0100", "0110", 1)
				count++
			}
			if strings.Contains(s, "0101") {
				s = strings.Replace(s, "0101", "0111", 1)
				count++
			}
		}
		fmt.Println(count)
	}
}
