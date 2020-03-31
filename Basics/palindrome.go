package main

import "fmt"

func main() {
	num := 0
	revNum := 0
	rem := 0
	var temp int

	num = 123456
	temp = num
	for temp != 0 {
		rem = temp % 10
		revNum = revNum*10 + rem
		temp /= 10
	}
	fmt.Println("Reverse number is ", revNum)
	if revNum == num {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}
