package main

import (
	"fmt"
)

// hasBit checks for set or unset bit
func hasBit(num int64, pos uint) bool {
	val := num & (1 << pos)
	return val > 0
}

// ChangeNBit - changes bit N from 1 to 0 and vice versa
func ChangeNBit(num int64, pos uint) int64 {
	// changes form #1-8 to #0-7
	pos--
	if hasBit(num, pos) {
		num &= ^(1 << pos)
	} else {
		num |= (1 << pos)
	}
	return num
}

func main() {
	num := int64(64)
	fmt.Printf("Number: %d\n in bytes: %08b\n", num, num)

	// changes bits from pos #1-8
	changedNum := ChangeNBit(num, 1)
	fmt.Printf("Number: %d\n in bytes:  %08b\n", changedNum, changedNum)
}
