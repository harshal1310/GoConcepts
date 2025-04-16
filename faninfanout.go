package main

import (
	"fmt"
	"sync"
)

// if we remove <- still code works but it act as bi drictional channel
func worker(id int, jobs <-chan int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing %d\n", id, job)
		resultChan <- job * job
	}

}
func main() {

	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup
	// fan out
	workers := 3

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send 5 jobs (fan-out input)
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // No more jobs

	//fan in

	go func() {
		wg.Wait()
		close(results) // Fan-in done
	}()

	//below code does same like below
	/*for {
		result, ok := <-results
		if !ok {
			break
		}
		fmt.Println("Result:", result)
	}*/

	// Fan-In: Collect results
	for result := range results {
		fmt.Println("Result:", result)
	}

}
