package main

import (
	"fmt"
	"strings"
)

func main() {

	str1 := "string1"
	str2 := "strinG1"

	str1 = strings.ToUpper(str1)
	str2 = strings.ToUpper(str2)

	fmt.Println("Result ", str1 == str2)
}
