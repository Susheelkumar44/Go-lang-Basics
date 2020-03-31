package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func count(data string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- data
		time.Sleep(time.Millisecond * 500)
	}
}
