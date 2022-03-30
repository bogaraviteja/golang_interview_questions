package main

import "fmt"

func evenOdd(n int) {
	if n&1 != 0 {
		fmt.Println("ODD")
	} else {
		fmt.Println("EVEN")
	}

}

func main() {
	evenOdd(25)
	evenOdd(30)
}
