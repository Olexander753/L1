package main

import "fmt"

// remove - the element i from slice a with saving order
func remove(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

// remove1 - the element i from slice a without saving order
func remove1(a []int, i int) []int {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)
	fmt.Println(remove(arr, 3))
	//fmt.Println(remove1(arr, 3))
}
