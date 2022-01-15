package main

import (
	"fmt"
	"strings"
)

// reverseWords reverse words
func reverseWords(str string) string {
	words := strings.Fields(str)

	left, right := 0, len(words)-1

	for left <= right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}

// reverseWords reverse words
func reverseWords1(str string) string {
	words := strings.Fields(str)

	var res strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		res.WriteString(words[i])
		res.WriteString(" ")
	}

	return strings.TrimSpace(res.String())
}

func main() {
	s := "snow dog sun"

	fmt.Println(s)
	fmt.Println(reverseWords(s))
	fmt.Println(reverseWords1(s))
}
