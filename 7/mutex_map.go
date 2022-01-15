package main

import (
	"fmt"
	"sync"
)

// Data - Wrapping for sync map
type Data struct {
	sync.RWMutex
	item map[string]string
}

// NewData - creates new Data struct
func NewData() *Data {
	return &Data{
		item: make(map[string]string),
	}
}

// Add new data by key
func (d *Data) Add(key, value string) {
	d.Lock()
	d.item[key] = value
	d.Unlock()
}

// Get data by key
func (d *Data) Get(key string) string {
	d.RLock()
	item := d.item[key]
	d.RUnlock()

	return item
}

// Delete data by key
func (d *Data) Delete(key string) {
	d.Lock()

	delete(d.item, key)

	d.Unlock()
}

func main() {
	data := NewData()

	var wg sync.WaitGroup

	// write data in data
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			data.Add(fmt.Sprintf("%d", i), fmt.Sprintf("%d", i*i))
		}
	}()

	// read from data
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			num := data.Get(fmt.Sprintf("%d", i))
			if len(num) < 1 {
				fmt.Println("Deleted - First")
			} else {
				fmt.Println(num, " - First")
			}
		}
	}()

	// read from data
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			num := data.Get(fmt.Sprintf("%d", i))
			if len(num) < 1 {
				fmt.Println("\t\tDeleted - Second")
			} else {
				fmt.Println("\t\t", num, " - Second")
			}
		}
	}()

	// delete some data
	data.Delete("1")
	data.Delete("2")

	wg.Wait()
}
