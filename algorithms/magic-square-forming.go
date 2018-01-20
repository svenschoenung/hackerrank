package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)

func isMagic(square []int) bool {
	if square[0]+square[1]+square[2] == 15 &&
		square[3]+square[4]+square[5] == 15 &&
		square[6]+square[7]+square[8] == 15 &&
		square[0]+square[3]+square[6] == 15 &&
		square[1]+square[4]+square[7] == 15 &&
		square[2]+square[5]+square[8] == 15 &&
		square[0]+square[4]+square[8] == 15 &&
		square[2]+square[4]+square[6] == 15 {
		return true
	}
	return false
}

// Heap's algorithm
func permutate(sq []int, n int, fn func([]int)) {
	if n == 1 {
		fn(sq)
		return
	}

	for i := 0; i < n-1; i++ {
		permutate(sq, n-1, fn)
		if n%2 == 0 {
			sq[i], sq[n-1] = sq[n-1], sq[i]
		} else {
			sq[0], sq[n-1] = sq[n-1], sq[0]
		}
	}
	permutate(sq, n-1, fn)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func dist(sq1 []int, sq2 []int) int {
	dist := 0
	for i := 0; i < 9; i++ {
		dist += abs(sq1[i] - sq2[i])
	}
	return dist
}

func main() {
	square := make([]int, 9)
	fmt.Scanf("%d %d %d\n", &square[0], &square[1], &square[2])
	fmt.Scanf("%d %d %d\n", &square[3], &square[4], &square[5])
	fmt.Scanf("%d %d %d\n", &square[6], &square[7], &square[8])

	minDist := MaxInt
	permutate([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, func(sq []int) {
		if isMagic(sq) {
			minDist = min(minDist, dist(square, sq))
		}
	})
	fmt.Printf("%d\n", minDist)
}
