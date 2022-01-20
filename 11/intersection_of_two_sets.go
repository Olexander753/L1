package main

import "fmt"

// intersection - intersects two slices
func intersection(a, b []int) []int {
	set1 := make(map[int]bool, len(a))
	set2 := make(map[int]bool, len(b))

	var res []int

	for _, val := range a {
		set1[val] = true
	}

	for _, val := range b {
		set2[val] = true
	}

	for key := range set2 {
		if set1[key] {
			res = append(res, key)
		}
	}

	return res
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	set2 := []int{4, 5, 6, 8, 1, 10, 11, 0}

	fmt.Println(intersection(set1, set2))
}
