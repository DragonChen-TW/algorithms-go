package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch4 ex1-4 Matrix Addition
// Date:	2022/11/30

// Textbook p85
func MatrixAddRecursive(A, B, C [][]int) {
	// fmt.Println(A, B, C)
	var n int = len(A)
	if n == 0 || len(A[0]) == 0 {
		return
	}
	if n == 1 {
		for i := range A[0] {
			C[0][i] = A[0][i] + B[0][i]
		}
		return
	}

	A11, A12, A21, A22 := MatrixPartition(A)
	// fmt.Println("A", A, A11, A12, A21, A22)
	B11, B12, B21, B22 := MatrixPartition(B)
	C11, C12, C21, C22 := MatrixPartition(C)

	MatrixAddRecursive(A11, B11, C11)
	MatrixAddRecursive(A12, B12, C12)
	MatrixAddRecursive(A21, B21, C21)
	MatrixAddRecursive(A22, B22, C22)
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

// modified from 1_matrix_multiply.go
func MatrixAdd(A, B [][]int) [][]int {
	var m, n int = len(A), len(A[0])
	var C [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		C[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}

	return C
}

func main() {
	fmt.Println(MatrixAdd(
		[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
		[][]int{{11, 12}, {13, 14}, {15, 16}, {17, 18}, {19, 20}},
	))
	fmt.Println(MatrixAdd(
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{11, 22, 33}, {44, 55, 66}, {77, 88, 99}},
	))

	n := 3
	var out [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		out[i] = make([]int, n)
	}
	MatrixAddRecursive(
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[][]int{{11, 22, 33}, {44, 55, 66}, {77, 88, 99}},
		out,
	)
	fmt.Println(out)
}
