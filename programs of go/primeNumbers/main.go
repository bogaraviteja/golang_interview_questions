package main

import (
	"fmt"
	"math"
)

func primeNumber(n1, n2 int) {
	if n1 < 2 || n2 < 2 {
		fmt.Println("Numbers must be greater than 2")
		return
	}
	for n1 <= n2 {
		isPrime := true
		for i := 2; i < int(math.Sqrt(float64(n1))); i++ {
			if n1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d", n1)
		}
		n1++
	}
	fmt.Println()
}

func main() {
	primeNumber(0, 25)
	primeNumber(2, 25)

}
