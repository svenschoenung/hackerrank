package main

import "fmt"

func main() {
	var size int
	fmt.Scanf("%d\n", &size)

	var score int
	fmt.Scanf("%d", &score)
	best := score
	worst := score

	brokeBest := 0
	brokeWorst := 0

	for i := 1; i < size; i++ {
		fmt.Scanf("%d", &score)
		if score > best {
			best = score
			brokeBest++
		}
		if score < worst {
			worst = score
			brokeWorst++
		}
	}

	fmt.Printf("%d %d", brokeBest, brokeWorst)
}
