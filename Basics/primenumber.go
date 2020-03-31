package main

import (
	"fmt"
	//"math/big"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	//Use of in-built function
	// i := big.NewInt(4)
	// isPrime := i.ProbablyPrime(1)
	// fmt.Println(isPrime)

	fmt.Println(isPrime(4))
}
