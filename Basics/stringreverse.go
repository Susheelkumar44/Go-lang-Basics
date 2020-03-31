package main

import (
	"fmt"
)

func main() {

	var revStr = make([]byte, 0)
	str := []byte("abcdef")
	for i := (len(str) - 1); i >= 0; i-- {
		revStr = append(revStr, str[i])
	}
	if string(str) == string(revStr) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}

	fmt.Println(string(revStr))
}
