package main

import "fmt"

func main() {
	var mat1 [2][2]int
	var mat2 [2][2]int
	var addmat [2][2]int

	fmt.Print("Enter the elements of matrix A : ")
	for i, j := range mat1 {
		for k := range j {
			fmt.Scan(&mat1[i][k])
		}
	}

	fmt.Print("Enter the elements of matrix B : ")
	for i, j := range mat2 {
		for k := range j {
			fmt.Scan(&mat2[i][k])
		}
	}

	fmt.Println("The Addition of Two matrices :")
	for i, row := range addmat {
		for j := range row {
			addmat[i][j] = mat1[i][j] + mat2[i][j]
			fmt.Print(addmat[i][j], "\t")
		}
		fmt.Println()
	}

}
