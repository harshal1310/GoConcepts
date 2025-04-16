package main

import (
	"fmt"
	"sync"
	"time"
)

func printName(name string, ch chan string) {
	ch <- name
}

func printNameWithGroup(name string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done
	ch <- name
}

func printMessage() {
	fmt.Println("Hello from Goroutine!")
}

func voidGo() {
	fmt.Println("in void go")
}

func voidGoWaitGroup(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("in void go wait group")

}

func withBuffer(ch chan string, wg *sync.WaitGroup) {
	ch <- "ch"
}

func process(job int) {
	fmt.Printf("Processing job %d\n", job)
	time.Sleep(time.Second)
	fmt.Printf("Finished job %d\n", job)
}

func main() {
	//name := go printName("ash")
	//fmt.Println(name)'

	//without waitgroup
	/*printMessage()
	time.Sleep(time.Second)

	ch := make(chan string)
	go printName("ash", ch)
	name := <-ch
	fmt.Println(name)

	for i := 0; i < 3; i++ {
		go voidGo()
	}
	time.Sleep(time.Second)
	*/

	//with waitgroup

	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan string)
	go printNameWithGroup("waitgroupname", ch, &wg)
	name := <-ch
	fmt.Println(name)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go voidGoWaitGroup(&wg)
	}

	chBuffer := make(chan string, 1)
	chBuffer <- "ch1" //removes value from channel
	//uncomment below line fto resolve deadlock issue
	//fmt.println("reading value from channel",<-channel)
	//x:<-chanBuffer
	chBuffer <- "ch2"

	wg.Wait() //it waits until we dont use defer wg.done

	//in worker pool we are  restricting goroutine count like numofjobs goruouines will be used after one gorouine is free it's resued also limiting channel with buffer

	/*
	   numJobs := 5
	   jobs := make(chan int, 2) // Buffered channel of size 2

	   // Producer: Sending jobs

	   	go func() {
	   		for i := 1; i <= numJobs; i++ {
	   			jobs <- i
	   			fmt.Println("Sent job", i)
	   		}
	   		close(jobs)
	   	}()

	   // Consumer: Processing jobs

	   	for job := range jobs {
	   		go process(job) // Creates a new goroutine for each job (NO worker pool)
	   	}

	   time.Sleep(3 * time.Second) // Wait for jobs to complete
	*/
}
