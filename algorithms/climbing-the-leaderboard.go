package main

import "fmt"

func rankHighScores(highScores []int, ranks []int) {
	previousHighScore := -1
	rank := 0
	for i, highScore := range highScores {
		if highScore != previousHighScore {
			rank++
		}
		ranks[i] = rank
		previousHighScore = highScore
	}
}

func findRank(highScores []int, ranks []int, score int, pos int) (int, int) {
	for i := pos; i >= 0; i-- {
		if highScores[i] > score {
			return ranks[i] + 1, i
		}
		if highScores[i] == score {
			return ranks[i], i
		}
	}
	return 1, 0
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)

	highScores := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &highScores[i])
	}

	ranks := make([]int, n)
	rankHighScores(highScores, ranks)

	var m int
	fmt.Scanf("%d\n", &m)

	pos := n - 1
	for i := 0; i < m; i++ {
		var score, rank int
		fmt.Scanf("%d", &score)
		rank, pos = findRank(highScores, ranks, score, pos)
		fmt.Printf("%d\n", rank)
	}
}
