package main

import "fmt"

// Author:    DragonChen https://github.com/dragonchen-tw/
// Title:    Part1 Ch2 ex1-5 Add Binary Number
// Date:    2022/10/14

func main() {
	// Procedure
	fmt.Println(`AddBinaryNumbers(A, B)
----------
buf = 0
for i = 0 to n-1
C[i] = A[i] + B[i] + buf
if C[i] > 1
	C[i] = C[i] / 2
	buf = C[i] % 2
else
	buf = 0
C[n] = buf
return C
----------`)

	// Program
	fmt.Println(AddBinaryNumbers([]byte{0, 1, 1}, []byte{1, 0, 1})) // {1, 1, 0, 1}

	// // Testing
	// var length int = 1000000000
	// var a, b []byte = make([]byte, length), make([]byte, length)
	// for i := 0; i < length; i++ {
	// 	a[i] = 1
	// 	b[i] = 1
	// }
	// AddBinaryNumbers(a, b)

	// // []bool version
	// var a, b []bool = make([]bool, length), make([]bool, length)
	// for i := 0; i < length; i++ {
	// 	a[i] = true
	// 	b[i] = true
	// }

	// Speed test with 1,000,000,000 1 billion
	// []byte 3.996 sec total
	// []bool 3.434 sec total
}

// []byte
func AddBinaryNumbers(A, B []byte) []byte {
	var C []byte = make([]byte, len(A)+1)
	var buf byte = 0
	for i := 0; i < len(A); i++ {
		C[i] = A[i] + B[i] + buf

		if C[i] > 1 {
			buf = C[i] / 2
			C[i] = C[i] % 2
		} else {
			buf = 0
		}
	}
	C[len(A)] = buf

	return C
}

// // []bool
// func AddBinaryNumbers(A, B []bool) []bool {
// 	var C []bool = make([]bool, len(A)+1)
// 	var buf bool = false
// 	for i := 0; i < len(A); i++ {
// 		if A[i] && B[i] {
// 			C[i] = buf
// 			buf = true
// 		} else if A[i] || B[i] {
// 			if buf {
// 				C[i] = false
// 			} else {
// 				C[i] = true
// 				buf = false
// 			}
// 		} else {
// 			C[i] = buf
// 			buf = false
// 		}
// 	}
// 	C[len(A)] = buf

// 	return C
// }
