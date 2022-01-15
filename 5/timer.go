package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Read - reads from channel
func Read(ch chan string) {
	counter := 1
	for data := range ch {
		fmt.Println("read data: ", data, " ", counter, " time(s)")
		counter++
	}
	fmt.Println("channel closed")
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Write 1 arg: time of work in ms")
		return
	}

	t, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You need to write a number")
		return
	}

	ch := make(chan string)
	go Read(ch)

	duration := time.Millisecond * time.Duration(t)
	timer := time.NewTimer(duration)

	for {
		select {
		case <-timer.C:
			close(ch)
			fmt.Println("time's out")
			time.Sleep(time.Millisecond * 100)
			return
		default:
			ch <- "Hello World!"
		}
		time.Sleep(time.Millisecond * 100)
	}
}
