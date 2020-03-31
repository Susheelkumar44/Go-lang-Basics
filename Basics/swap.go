package main

import "fmt"

func main() {
	a, b := 10, 40
	fmt.Println("a, b", a, b)
	a, b = b, a
	fmt.Println("reverse of a, b", a, b)
}
