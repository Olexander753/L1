package main

import (
	"fmt"
	"time"
)

// Sleep - pause a thread on time t
func Sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	start := time.Now()
	fmt.Println("Program started")

	fmt.Println("Program will fall asleep for 2 seconds")
	Sleep(time.Second * 2)

	fmt.Println("Program fineshed with time: ", time.Since(start))
}
