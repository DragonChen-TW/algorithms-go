package main

import "fmt"

// Author:	DragonChen https://github.com/dragonchen-tw/
// Title:	Part1 Ch2 3-8 Find Two Sum
// Date:	2022/11/24

// Textbook p45
func FindTwoSum(nums []int, target int) bool {
	MergeSort(nums, 0, len(nums)-1)
	for i, n := range nums {
		if idx := BinarySearch(nums, target-n, i); idx != -1 {
			return true
		}
	}
	return false
}

// modified from 2.3-6
func BinarySearch(nums []int, target int, idx int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			if mid == idx {
				l++
				continue
			}
			return mid
		}
	}

	// return the left-most position instead of -1 when target is not found
	return -1
}

// Reference from 2.3-1
func MergeSort(nums []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2
	MergeSort(nums, p, q)
	MergeSort(nums, q+1, r)
	Merge(nums, p, q, r)
}
func Merge(nums []int, p, q, r int) {
	lenL := q - p + 1
	lenR := r - q
	var L []int = make([]int, lenL)
	var R []int = make([]int, lenR)
	copy(L, nums[p:q+1])
	copy(R, nums[q+1:r+1])

	var i, j, k int = 0, 0, p
	for i < lenL && j < lenR {
		if L[i] <= R[j] {
			nums[k] = L[i]
			i++
		} else {
			nums[k] = R[j]
			j++
		}
		k++
	}

	for i < lenL {
		nums[k] = L[i]
		i++
		k++
	}
	for j < lenR {
		nums[k] = R[j]
		j++
		k++
	}
}

func main() {
	fmt.Println(FindTwoSum([]int{3, 1, 7, 8, 4, 3}, 10)) // true
	fmt.Println(FindTwoSum([]int{3, 1, 7, 8, 4, 3}, 11)) // true
	fmt.Println(FindTwoSum([]int{3, 1, 7, 8, 4, 3}, 12)) // true
	fmt.Println(FindTwoSum([]int{3, 1, 7, 8, 4, 3}, 13)) // false
}
