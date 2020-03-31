package main

import "fmt"

func main() {
	n := 10

	t1 := 0
	t2 := 1
	next := 0

	for i := 1; i <= n; i++ {
		next = t1 + t2
		t1 = t2
		t2 = next
		fmt.Println("next", next)
	}
	fmt.Println("\n")

}
