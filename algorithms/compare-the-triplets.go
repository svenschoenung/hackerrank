package main

import "fmt"

func readScores() [3]int {
	var scores [3]int
	for i := 0; i < 3; i++ {
		fmt.Scanf("%v", &scores[i])
	}
	return scores
}

func main() {
	alizeScores := readScores()
	bobScores := readScores()

	alizePoints := 0
	bobPoints := 0

	for i := 0; i < 3; i++ {
		if alizeScores[i] > bobScores[i] {
			alizePoints++
		} else if alizeScores[i] < bobScores[i] {
			bobPoints++
		}
	}

	fmt.Printf("%d %d\n", alizePoints, bobPoints)
}
