package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch4 1 Matrix Multiply
// Date:	2022/11/29

// Textbook p81 normal version
func MatrixMultiply(A, B [][]int) [][]int {
	var n int = len(A)
	var C [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}

// Textbook p84
// already fit the ex4.1-1
func MatrixMultiplyRecursive(A, B, C [][]int) {
	var n int = len(A)
	if n == 0 {
		return
	}
	if n == 1 {
		C[0][0] += A[0][0] * B[0][0]
		return
	}

	A11, A12, A21, A22 := MatrixPartition(A)
	B11, B12, B21, B22 := MatrixPartition(B)
	C11, C12, C21, C22 := MatrixPartition(C)

	MatrixMultiplyRecursive(A11, B11, C11)
	MatrixMultiplyRecursive(A12, B21, C11)
	MatrixMultiplyRecursive(A11, B12, C12)
	MatrixMultiplyRecursive(A12, B22, C12)
	MatrixMultiplyRecursive(A21, B11, C21)
	MatrixMultiplyRecursive(A22, B21, C21)
	MatrixMultiplyRecursive(A21, B12, C22)
	MatrixMultiplyRecursive(A22, B22, C22)
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

func main() {
	fmt.Println(MatrixMultiply(
		[][]int{{1, 2}, {3, 4}},
		[][]int{{5, 6}, {7, 8}},
	))

	var out [][]int = make([][]int, 2)
	for i := 0; i < 2; i++ {
		out[i] = make([]int, 2)
	}
	MatrixMultiplyRecursive([][]int{{1, 2}, {3, 4}}, [][]int{{5, 6}, {7, 8}}, out)
	fmt.Println(out)
}
