package main

import "fmt"

func reverseNumber(n int) int {
	r := 0
	for n > 0 {
		remainder := n % 10
		r = (r * 10) + remainder
		n /= 10
	}
	return r
}

func main() {
	fmt.Println(reverseNumber(123456789))
}
