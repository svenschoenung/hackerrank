package main

import "fmt"
import "strconv"

func isBeautiful(minL int, num int64, runes []rune) bool {
	for l := minL; l <= minL+1 && l <= len(runes); l++ {
		if runes[:l][0] != '0' {
			num2, _ := strconv.ParseInt(string(runes[:l]), 10, 64)
			rest := runes[l:]
			if num2-num == 1 {
				if len(rest) == 0 {
					return true
				}
				return isBeautiful(l, num2, rest)
			}
		}
	}
	return false
}

func checkBeautiful(s string) {
	runes := []rune(s)
	for l := 1; l < len(runes); l++ {
		if l == 1 || runes[:l][0] != '0' {
			num, _ := strconv.ParseInt(string(runes[:l]), 10, 64)
			rest := runes[l:]
			if isBeautiful(l, num, rest) {
				fmt.Printf("YES %d\n", num)
				return
			}
		}
	}
	fmt.Printf("NO\n")
}

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%v\n", &s)

		checkBeautiful(s)
	}
}
