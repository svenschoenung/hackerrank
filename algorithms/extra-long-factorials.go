package main

import "fmt"
import "math/big"

var ZERO *big.Int = big.NewInt(0)
var ONE *big.Int = big.NewInt(1)

func main() {
	var n int64
	fmt.Scanf("%d\n", &n)

	var bigN *big.Int = big.NewInt(n)
	var fact *big.Int = big.NewInt(1)

	for bigN.Cmp(ZERO) > 0 {
		fact.Mul(fact, bigN)
		bigN.Sub(bigN, ONE)
	}
	fmt.Printf("%v\n", fact)
}
