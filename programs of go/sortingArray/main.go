package main

import "fmt"

func binarySearch(arr []int, key int) int {
	high := len(arr) - 1
	low := 0
	var mid int
	for low <= high {
		mid = (high + low) / 2
		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 4, 6, 8, 9, 10}, 11))
	fmt.Println(binarySearch([]int{1, 4, 6, 8, 9, 10}, 8))
	fmt.Println(binarySearch([]int{1, 4, 6, 8, 9, 10}, 10))
}
