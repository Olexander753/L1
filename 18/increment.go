package main

import (
	"fmt"
	"sync"
)

// Counter - counter with mutex for parallel counting
type Counter struct {
	sync.Mutex
	counter int
}

// NewCounter - constructor for counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc - incrementing counter by 1
func (c *Counter) Inc() {
	c.Lock()
	c.counter++
	c.Unlock()
}

// String - prints value of counter
func (c *Counter) String() string {
	return fmt.Sprint(c.counter)
}

func main() {

	c := NewCounter()
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Println(c)
}
