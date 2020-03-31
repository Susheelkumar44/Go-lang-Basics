package main

import "fmt"

func bubbleSort(toSort []int) {
	length := len(toSort)

	if length < 2 {
		return 
	}

	for i:=0; i< length; i++ {
		for j:=length-1; j>=i+1; j-- {
			if toSort[j] < toSort[j-1] {
				toSort[j], toSort[j-1] = toSort[j-1], toSort[j]
			}
		}
	}
}

func main() {
	unsorted := []int{9,6,44,2,0,8,3,98,198, 44}

	fmt.Println("Before Sorting: ", unsorted)

	bubbleSort(unsorted)

	fmt.Println("After Sorting: ", unsorted)
}