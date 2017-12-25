package main

import "fmt"

func roundGrade(grade int) int {
	if grade < 38 {
		return grade
	}
	if grade%5 >= 3 {
		return grade + (5 - grade%5)
	}
	return grade
}

func main() {
	var count int
	fmt.Scanf("%d\n", &count)

	output := ""
	for i := 0; i < count; i++ {
		var grade int
		fmt.Scanf("%d\n", &grade)
		output += fmt.Sprintf("%d\n", roundGrade(grade))
	}
	fmt.Print(output)
}
