package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Test 2-dim Matrix
// Date:	2022/11/29

func main() {
	// matrix := CreateMatrix(4)

	// fmt.Println(matrix)
	// for m := range matrix {
	// 	fmt.Println(matrix[m])
	// }
	fmt.Println(MatrixPartition(CreateMatrix(2)))
	fmt.Println(MatrixPartition(CreateMatrix(3)))
	fmt.Println(MatrixPartition(CreateMatrix(4)))

	m := CreateMatrix(3)
	m1, m2, m3, m4 := MatrixPartition(m)
	m1[0][0] = -1
	m2[0] = []int{11, 22}
	m3[1][0] = 99
	m4[0][1] = 18
	fmt.Println(m1, m2, m3, m4)
	fmt.Println(m)
}

func MatrixPartition(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(matrix)
	n1, n2 := n/2, n-n/2
	var m11, m12, m21, m22 [][]int
	m11 = make([][]int, n1) // n1 x n1
	m12 = make([][]int, n1) // n1 x n2
	m21 = make([][]int, n2) // n2 x n1
	m22 = make([][]int, n2) // n2 x n2
	for i := 0; i < n1; i++ {
		m11[i] = matrix[i][:n1]
		m12[i] = matrix[i][n1:]
	}
	for i := 0; i < n2; i++ {
		m21[i] = matrix[n1+i][:n1]
		m22[i] = matrix[n1+i][n1:]
	}
	return m11, m12, m21, m22
}

func CreateMatrix(n int) [][]int {
	var matrix [][]int
	matrix = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = i*n + j
		}
	}
	return matrix
}
