package main

import (
	"fmt"
)

// SumSquares - parallel calculation of the sum of squares of numbers
func SumSquares(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var sum int
	ch := make(chan int, len(nums))

	go func() {
		for i := 0; i < len(nums); i++ {
			go func(num int) {

				square := num * num
				ch <- square
			}(nums[i])
		}
	}()

	for i := 0; i < len(nums); i++ {
		sum += <-ch
	}

	close(ch)
	return sum
}

//  SumSquares2 - second version with 2 goroutines: main and child
func SumSquares2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	var sum int
	ch := make(chan int, len(nums))

	go func(ch chan int) {
		for _, v := range nums {
			ch <- v * v
		}
		close(ch)
	}(ch)

	for {
		val, ok := <-ch
		if !ok {
			break
		}
		sum += val
	}

	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	//sum := SumSquares(numbers)
	sum := SumSquares2(numbers)

	fmt.Printf("Sum of squares is equal %d\n", sum)
}
