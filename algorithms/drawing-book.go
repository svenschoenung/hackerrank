package main

import "fmt"

func counts(arr []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range arr {
		counts[num]++
	}
	return counts
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func keys(m map[int]int) []int {
	keys := make([]int, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	counts := counts(arr)
	nums := keys(counts)

	result := 0
	for i := 0; i < max(len(nums)-1, 1); i++ {
		num1 := nums[i]
		count1 := counts[num1]
		result = max(result, count1)
		for j := i + 1; j < len(nums); j++ {
			num2 := nums[j]
			count2 := counts[num2]
			if abs(num1-num2) <= 1 {
				result = max(result, count1+count2)
			}
		}
	}
	fmt.Printf("%d", result)
}
