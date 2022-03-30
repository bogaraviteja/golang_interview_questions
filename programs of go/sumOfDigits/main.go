package main

import "fmt"

func findDigitSum(n int) int {
	r := 0
	for n > 0 {
		r += n % 10
		n /= 10
	}
	return r
}

func main() {
	fmt.Println(findDigitSum(168))
	fmt.Println(findDigitSum(576))
	fmt.Println(findDigitSum(12345))
}
