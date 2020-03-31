package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{9,5,3,2,8,4}

	sort.Ints(arr)
	fmt.Println("Second smallest element is ",arr[1])
	fmt.Println("Second Largest element is ", arr[(len(arr)-2)])
}