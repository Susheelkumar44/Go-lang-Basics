package main

import "fmt"

func main() {
	array := [6]int{10,44,77,2,90,6}
	value := array[0]

	for _, val := range array {
		if val < value {
			value=val
		} 
	}
	fmt.Println("Smallest number is", value)
}