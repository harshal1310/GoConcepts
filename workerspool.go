package main

import (
	"fmt"
	"sync"
	"time"
)

func process(job int) {
	fmt.Printf("Processing job %d\n", job)
	time.Sleep(time.Second)
	fmt.Printf("Finished job %d\n", job)
}
func processWithWaitGroup(job int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done
	fmt.Printf("Processing job %d\n", job)
	time.Sleep(time.Second)
	fmt.Printf("Finished job %d\n", job)
}
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		process(job)
	}

}

func main() {
	// Record the start time
	startTime := time.Now()
	numJobs := 10
	numWorkers := 3
	var wg sync.WaitGroup
	jobs := make(chan int, 5)

	// Send jobs to the jobs channel
	go func() {
		for i := 1; i <= numJobs; i++ {
			fmt.Printf("Sending job %d\n", i)
			jobs <- i
		}
		close(jobs) // Close the jobs channel to signal workers that no more jobs will be sent
	}()

	for job := 0; job < numWorkers; job++ {
		wg.Add(1)
		go worker(job, jobs, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()
	elapsedTime := time.Since(startTime)
	fmt.Printf("All jobs processed in %s\n", elapsedTime)
	/*

		// Record the start time
		startTime := time.Now()

		numJobs := 10
		jobs := make(chan int, 5) // Buffered channel to hold jobs
		var wg sync.WaitGroup     // WaitGroup to wait for all jobs to finish

		// Send jobs to the jobs channel
		go func() {
			for i := 1; i <= numJobs; i++ {
				fmt.Printf("Sending job %d\n", i)
				jobs <- i
			}
			close(jobs) // Close the jobs channel to signal no more jobs will be sent
		}()

		// Process jobs without a worker pool
		for job := range jobs {
			wg.Add(1) // Increment the WaitGroup counter for each job
			go func(job int) {
				defer wg.Done() // Mark this goroutine as done
				process(job)
			}(job)
		}

		// Wait for all jobs to finish
		wg.Wait()

		// Calculate and print the total time taken
		elapsedTime := time.Since(startTime)
		fmt.Printf("All jobs processed in %s\n", elapsedTime)
	*/
}
