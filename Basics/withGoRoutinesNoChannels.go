package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func count(data string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, data)
		time.Sleep(time.Millisecond * 500)
	}
}
