package main

import (
	"fmt"
	"sync"
)

// Squares - parallel calculation of the area of array with waitgroup
func Squares(nums []int) []int {
	if len(nums) < 1 {
		return []int{}
	}

	var wg sync.WaitGroup

	squares := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()

			square := nums[idx] * nums[idx]

			squares[idx] = square
		}(i)
	}

	wg.Wait()
	return squares
}

// Squares - concurrency compute squre of array with channel
func SquaresWithChan(nums []int) []int {
	if len(nums) < 1 {
		return []int{}
	}

	ch := make(chan int)
	squares := make([]int, 0, len(nums))

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
		squares = append(squares, val)
	}

	return squares
}

func main() {

	numbers := [5]int{2, 4, 6, 8, 10}

	//squares := Squares(numbers[:])
	squares := SquaresWithChan(numbers[:])

	for i := range numbers {
		fmt.Printf("square of %d is %d\n", numbers[i], squares[i])
	}
}
