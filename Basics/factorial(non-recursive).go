package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Println("Enter the number ")
	fmt.Scanf("%d", &n)
	a:=n

	//var i = 0
	for i := (n - 1); i >= 1; i-- {
		n = n * i
	}

	if n == 0 {
		fmt.Printf("Factorial of %v is: %+v",a, 1)
	} else {
		fmt.Printf("Factorial of %+v is: %+v",a, n)
	}
}
