package main

import "fmt"

func main() {
	str1 := "abcdef"
	var answer bool

	for i := 0; i < len(str1)/2; i++ {
		if str1[i] != str1[len(str1)-i-1] {
			answer = false
		}
		answer = true
	}

	if answer {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}
