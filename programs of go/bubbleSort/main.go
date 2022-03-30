package main

import "fmt"

func checkSortedArray(arr []int) {
	sortedArray := true
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				sortedArray = false
				break
			}
		}
	}
	if sortedArray {
		fmt.Println("Given array is already sorted.")
	} else {
		fmt.Println("Given array is not sorted.")
	}
}

func main() {
	checkSortedArray([]int{1, 3, 5, 6, 7, 8})
	checkSortedArray([]int{1, 3, 5, 9, 4, 2})
	checkSortedArray([]int{9, 7, 4, 2, 1, -1})
}
