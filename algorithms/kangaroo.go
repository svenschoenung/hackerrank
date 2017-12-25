package main

import "fmt"
import "math"

func printAnswer(answer bool) {
	if answer {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

func main() {
	var x1, v1 int
	fmt.Scanf("%d %d", &x1, &v1)

	var x2, v2 int
	fmt.Scanf("%d %d", &x2, &v2)

	if v1 == v2 {
		printAnswer(x1 == x2)
	} else {
		//x1 + v1 * n = x2 + v2 * n
		//solve for n
		n := (float64(x2) - float64(x1)) / (float64(v1) - float64(v2))
		printAnswer(n >= 0 && math.Floor(n) == math.Ceil(n))
	}

}
