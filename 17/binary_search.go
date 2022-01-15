package main

import "fmt"

// binarySearch - search item in sorted array
func binarySearch(nums []int, target int) int {
	var min int
	max := len(nums) - 1

	for min <= max {
		mid := (max + min) / 2 // min + ((max - min) / 2) to fix overflow bug
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] > target:
			max = mid - 1
		case nums[mid] < target:
			min = mid + 1
		}
	}
	return -1
}

func main() {
	sortedArr := []int{-1, 0, 3, 5, 9, 12}

	fmt.Println(binarySearch(sortedArr, 9))
}
