package main

import "fmt"

// reverseString - reverse bytes in the string
func reverseString(s string) string {
	runes := []rune(s)
	left, right := 0, len(runes)-1

	for left <= right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func main() {

	a := "& главрыба ∧"
	fmt.Println(a)
	fmt.Println(reverseString(a))
}
