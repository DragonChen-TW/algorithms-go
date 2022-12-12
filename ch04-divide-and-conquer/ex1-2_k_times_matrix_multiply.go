package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch4 ex1-2 K times Matrix Multiply
// Date:	2022/11/29

// Multiplication of matrix (shape of kn, n) and matrix (shape of n, kn)
// We can split the matrix (shape of kn, n and n, kn) into k matrixs (shape of n, n).
// Then, compute outputed matrix by calling MatrixMultiplyRecursive k*k times.
// So, O(n) = n*2
//
// In the other hand, multiplication of matrix (shape of n, kn) and matrix (shape of kn, n)
// Has less mulication and one more addition.
// So, O(n) = n + O(addition) = n
// It faster k times than first mulplication

// Textbook p84
func KTimesMatrixMultiplyRecursive(A, B, C [][]int) {
	k, n := len(A), len(A[0])
	if k%n != 0 {
		fmt.Printf("Dimension of A (%d, %d) is wrong, first number shoulbe be k times of second one", k, n)
		return
	}
	k = k / n

	// k*n, n x n, k*n
	var smA, smB, smC [][]int
	smA = make([][]int, n)
	smB = make([][]int, n)
	smC = make([][]int, n)
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			for k := 0; k < n; k++ {
				smA[k] = A[i*n+k]
				smB[k] = B[k][j*n : (j+1)*n]
				smC[k] = C[i*n+k][j*n : (j+1)*n]
			}
			MatrixMultiplyRecursive(smA, smB, smC)
		}
	}
}

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

// modified from 1_matrix_multiply.go
func MatrixMultiply(A, B [][]int) [][]int {
	var m, n, o int = len(A), len(A[0]), len(B[0])
	var C [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		C[i] = make([]int, o)
	}

	// 1,2 x 2,3 => 1,3
	for i := 0; i < m; i++ {
		for j := 0; j < o; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C
}

func main() {
	k := 3
	n := 2

	fmt.Println(MatrixMultiply(
		[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
		[][]int{{11, 12, 13, 14, 15, 16}, {17, 18, 19, 20, 21, 22}},
	))
	fmt.Println(MatrixMultiply(
		[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}},
		[][]int{{13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}},
	))

	var out2 [][]int = make([][]int, k*n)
	for i := 0; i < k*n; i++ {
		out2[i] = make([]int, k*n)
	}
	KTimesMatrixMultiplyRecursive(
		[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}, {11, 12}},
		[][]int{{13, 14, 15, 16, 17, 18}, {19, 20, 21, 22, 23, 24}},
		out2,
	)
	fmt.Println(out2)
}
