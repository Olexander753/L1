package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	alpha := [26]int{}
	s = strings.ToLower(s)

	for _, v := range s {
		alpha[v-'a']++
	}

	for _, val := range alpha {
		if val > 1 {
			return false
		}
	}
	return true
}

func main() {
	strs := []string{"abcd", "abCdefAaf", "aabcd", "aAbcd"}

	for i := range strs {
		fmt.Printf("%s is %v\n", strs[i], isUnique(strs[i]))
	}
}
