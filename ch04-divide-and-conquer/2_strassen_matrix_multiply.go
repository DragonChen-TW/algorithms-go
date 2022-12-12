package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch4 2 Strassen Matrix Multiply
// Date:	2022/12/01

type OpType int

const (
	PLUS OpType = iota
	MINUS
)

// Textbook p84
// already fit the ex4.1-1 and ex4.2-2
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

	// make S matrixs
	S1 := MatrixPlusMinus(B12, B22, MINUS)
	S2 := MatrixPlusMinus(A11, A12, PLUS)
	S3 := MatrixPlusMinus(A21, A22, PLUS)
	S4 := MatrixPlusMinus(B21, B11, MINUS)
	S5 := MatrixPlusMinus(A11, A22, PLUS)
	S6 := MatrixPlusMinus(B11, B22, PLUS)
	S7 := MatrixPlusMinus(A12, A22, MINUS)
	S8 := MatrixPlusMinus(B21, B22, PLUS)
	S9 := MatrixPlusMinus(A11, A21, MINUS)
	S10 := MatrixPlusMinus(B11, B12, PLUS)

	P1 := CreateMatrix(n / 2)
	P2 := CreateMatrix(n / 2)
	P3 := CreateMatrix(n / 2)
	P4 := CreateMatrix(n / 2)
	P5 := CreateMatrix(n / 2)
	P6 := CreateMatrix(n / 2)
	P7 := CreateMatrix(n / 2)
	MatrixMultiplyRecursive(A11, S1, P1)
	MatrixMultiplyRecursive(S2, B22, P2)
	MatrixMultiplyRecursive(S3, B11, P3)
	MatrixMultiplyRecursive(A22, S4, P4)
	MatrixMultiplyRecursive(S5, S6, P5)
	MatrixMultiplyRecursive(S7, S8, P6)
	MatrixMultiplyRecursive(S9, S10, P7)

	MatrixPlusMinusRef(
		C11,
		MatrixPlusMinus(MatrixPlusMinus(MatrixPlusMinus(P5, P4, PLUS), P2, MINUS), P6, PLUS),
		PLUS,
	)
	MatrixPlusMinusRef(
		C12,
		MatrixPlusMinus(P1, P2, PLUS),
		PLUS,
	)
	MatrixPlusMinusRef(
		C21,
		MatrixPlusMinus(P3, P4, PLUS),
		PLUS,
	)
	MatrixPlusMinusRef(
		C22,
		MatrixPlusMinus(MatrixPlusMinus(MatrixPlusMinus(P5, P1, PLUS), P3, MINUS), P7, MINUS),
		PLUS,
	)
}

func CreateMatrix(n int) [][]int {
	var matrix [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	return matrix
}

func MatrixPartition(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(matrix)
	// n should be exact power of 2
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

// ref from 1_matrix_multiply.go
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

// modified from 1_matrix_multiply.go
func MatrixPlusMinus(A, B [][]int, op OpType) [][]int {
	var m, n int = len(A), len(A[0])
	var C [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		C[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if op == PLUS {
				C[i][j] = A[i][j] + B[i][j]
			} else if op == MINUS {
				C[i][j] = A[i][j] - B[i][j]
			}
		}
	}

	return C
}

// Slice version of MatrixPlusMinus
func MatrixPlusMinusRef(A, B [][]int, op OpType) {
	var m, n int = len(A), len(A[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if op == PLUS {
				A[i][j] += B[i][j]
			} else if op == MINUS {
				A[i][j] -= B[i][j]
			}
		}
	}
}

func main() {
	a := [][]int{
		{1, 2, 9, 3},
		{3, 4, 5, 7},
		{10, 3, 9, 1},
		{4, 7, 5, 9},
	}
	b := [][]int{
		{5, 6, 4, 1},
		{7, 8, 9, 12},
		{3, 2, 1, 0},
		{5, 6, 2, 10},
	}

	fmt.Println(MatrixMultiply(a, b))

	n := len(a)
	var out [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		out[i] = make([]int, n)
	}
	MatrixMultiplyRecursive(a, b, out)
	fmt.Println(out)

	a2 := [][]int{
		{1, 3},
		{7, 5},
	}
	b2 := [][]int{
		{6, 8},
		{4, 2},
	}
	n2 := len(a2)
	var out2 [][]int = make([][]int, n2)
	for i := 0; i < n2; i++ {
		out2[i] = make([]int, n2)
	}
	MatrixMultiplyRecursive(a2, b2, out2)
	fmt.Println(out2)
}
