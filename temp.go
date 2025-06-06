package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create a shared array
	data := make([]int, 5) // Array of size 5
	var wg sync.WaitGroup

	// Mutex for controlling access to the array
	var mu sync.Mutex

	// Writer goroutines
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Lock before modifying the shared array
			mu.Lock()
			data[i] = i * 10
			fmt.Printf("Writer %d wrote %d to data[%d]\n", i, i*10, i)
			mu.Unlock()
		}(i)
	}

	// Reader goroutine
	/*wg.Add(1)
	go func() {
		defer wg.Done()
		// Lock before reading the shared array
		mu.Lock()
		fmt.Println("Reader is reading data:")
		for idx, value := range data {
			fmt.Printf("data[%d] = %d\n", idx, value)
		}
		mu.Unlock()
	}()*/

	// Reader goroutine - Only start reading once all writers have finished
	go func() {
		wg.Wait() // Wait for all writers to finish
		// Lock before reading the shared array
		mu.Lock()
		fmt.Println("Reader is reading data:")
		for idx, value := range data {
			fmt.Printf("data[%d] = %d\n", idx, value)
		}
		mu.Unlock()
	}()

	// Wait for all goroutines to finish
	//wg.Wait()

	fmt.Println("Program finished.")
}
