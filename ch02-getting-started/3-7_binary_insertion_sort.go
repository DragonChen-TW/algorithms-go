package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 3-7 Insertion Sort with Binary Search
// Date:	2022/11/24

// Textbook p45
func InsertionSort(nums []int, n int) {
	// i, j change
	for i := 1; i < n; i++ {
		key := nums[i]

		// use Binary Search to locate the right posision in nums[0:i-1]
		idx := BinarySearch(nums[:i], key)

		// shift the sequence of idx to i - 1
		if idx < i-1 {
			copy(nums[idx+1:i+1], nums[idx:i])
		} else {
			nums[i] = nums[i-1]
		}

		nums[idx] = key
	}
}

// modified from 2.3-6
func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	// return the left-most position instead of -1 when target is not found
	return l
}

func main() {
	var a []int = []int{5, 2, 4, 6, 1, 3}
	InsertionSort(a, 6)
	fmt.Println(a)
}
