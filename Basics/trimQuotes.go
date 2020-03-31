package main

import (
	"fmt"
)

func trimQuotes(s string) string {
    if len(s) >= 2 {
        if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
            return s[1 : len(s)-1]
        }
    }
    return s
}

func main() {
	str1 := `"hello world"`
	fmt.Println("Before trimming quotes ", str1)
	str2 := trimQuotes(str1)
	fmt.Println("After trimming quotes ", str2)
}