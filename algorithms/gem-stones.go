package main

import "fmt"
import "strings"

func main() {
	var n int
	fmt.Scanf("%v\n", &n)

	gemstones := map[rune]bool{}
	for gemstone := 'a'; gemstone <= 'z'; gemstone++ {
		gemstones[gemstone] = true
	}

	for i := 0; i < n; i++ {

		var rock string
		fmt.Scanf("%v\n", &rock)

		for gemstone := 'a'; gemstone <= 'z'; gemstone++ {
			if !strings.Contains(rock, string(gemstone)) {
				gemstones[gemstone] = false
			}
		}
	}

	count := 0
	for _, foundInAll := range gemstones {
		if foundInAll {
			count++
		}
	}
	fmt.Println(count)
}
