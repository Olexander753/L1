package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// Workers - workers structure
type Workers struct {
	wg     sync.WaitGroup
	ch     chan interface{}
	cancel context.CancelFunc
	amount int
}

// NewWorkers - creates a new workers instance
func NewWorkers(amount int, cancel context.CancelFunc) *Workers {
	return &Workers{
		amount: amount,
		cancel: cancel,
		ch:     make(chan interface{}, amount),
	}
}

// Start workers
func (w *Workers) Start(ctx context.Context) {
	for i := 0; i < w.amount; i++ {
		w.wg.Add(1)
		go func(workerId int) {
			defer w.wg.Done()

			for {
				select {
				case data := <-w.ch:
					fmt.Printf("Worker %d read data from channel: Type %T, Value %v\n", workerId, data, data)
				case <-ctx.Done():
					fmt.Println("Worker ", workerId, " stopped")
					return
				}
			}
		}(i)
	}
}

// Stop all workers
func (w *Workers) Stop() {
	w.cancel()
	w.wg.Wait()
	close(w.ch)
	fmt.Println("All workers have been stopped")
}

// SendData - sends data via channel to workers
func (w *Workers) SendData(data interface{}) {
	w.ch <- data
}

func main() {

	var (
		numOfWorkers int
		err          error
	)

	// checking for arguments
	if len(os.Args) == 2 {
		numOfWorkers, err = strconv.Atoi(os.Args[1])
		if err != nil || numOfWorkers == 0 {
			fmt.Println("You need to write a number starting from 1")
			return
		}
	} else {
		fmt.Println("You need to specify 1 argument - amount of workers")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	workers := NewWorkers(numOfWorkers, cancel)
	workers.Start(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		num := generator.Intn(1000)
		chars := 'a' + rune(generator.Intn('z'-'a'+1)) // 'a' <= chars <= 'z'

		select {
		case signals := <-sigs:
			workers.Stop()
			fmt.Printf("Programm has been stopped by signal %d\n", signals)
			return
		default:
			workers.SendData(num)
			workers.SendData(string(chars))
		}

		time.Sleep(time.Millisecond * 100)
	}
}
