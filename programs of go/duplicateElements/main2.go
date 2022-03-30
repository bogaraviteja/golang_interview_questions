package main

import "fmt"

func duplicateInArray(arr []int, r int) int {
	xor := 0
	for i := 0; i < len(arr); i++ {
		xor ^= arr[i]
	}
	for j := 1; j <= r-1; j++ {
		xor ^= j
	}
	return xor
}
func main() {
	fmt.Println(duplicateInArray([]int{1, 2, 3, 4, 4}, 5))
	fmt.Println(duplicateInArray([]int{4, 5, 6, 7, 7}, 5))
}
