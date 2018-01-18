package main

import "fmt"

func main() {
	var steps int
	fmt.Scanf("%v\n", &steps)

	var input string
	fmt.Scanf("%s\n", &input)

	var level = 0
	var previousLevel = 0
	var valleys = 0
	for i := 0; i < steps; i++ {
		previousLevel = level
		step := input[i]
		if step == 'U' {
			level++
		} else if step == 'D' {
			level--
		}
		if previousLevel < 0 && level == 0 {
			valleys++
		}
	}
	fmt.Println(valleys)
}
