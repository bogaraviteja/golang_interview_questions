package main

import "fmt"

func main() {
	var rows, columns, i, j int

	var mat1 [10][10]int
	var mat2 [10][10]int
	var addMat [10][10]int

	fmt.Print("Enter the size of matrices: ")
	fmt.Scan(&rows, &columns)

	fmt.Print("Enter the elements of first matrix: ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&mat1[i][j])
		}
	}

	fmt.Print("Enter the elements of second matrix: ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Scan(&mat2[i][j])
		}
	}

	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			addMat[i][j] = mat1[i][j] + mat2[i][j]
		}
	}

	fmt.Println("The Addition of two matrices: ")
	for i = 0; i < rows; i++ {
		for j = 0; j < columns; j++ {
			fmt.Print(addMat[i][j], "\t")
		}
		fmt.Println()
	}

}
