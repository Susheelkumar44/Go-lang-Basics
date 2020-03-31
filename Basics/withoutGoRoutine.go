package main

import (
	"fmt"
	"time"
)

func main() {
	count("sheep")
	count("fish")
}

func count(data string) {
	for i := 0; true; i++ {
		fmt.Println(i, data)
		time.Sleep(time.Millisecond * 500)
	}
}
