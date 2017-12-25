package main

import "fmt"

func readSet(size int) []int {
	var set = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &set[i])
	}
	return set
}

func findFactors(num int) map[int]bool{
	factors := make(map[int]bool)
	for i := 1; i <= num / 2; i++ {
		if (num % i == 0) {
			factors[i] = true
		}
	}
	factors[num] = true;
	return factors
}

func intersection(set1 map[int]bool, set2 map[int]bool) {
	for k, _ := range set1 {
		if (!set2[k]) {
			delete(set1, k)
		}
	}
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)

	setA := readSet(n)
	setB := readSet(m)

	if (m <= 0 || n <= 0) {
		return
	}

	factors := findFactors(setB[0])
	for i := 1; i < m; i++ {
		intersection(factors, findFactors(setB[i]))
	}

	count := 0
	for factor, _ := range factors {
		between := true
		for _, val := range setA {
			if (factor % val != 0) {
				between = false
				break
			}
		}
		if between {
			count++
		}
	}
	fmt.Println(count)
}
