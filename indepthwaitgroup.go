package main

import (
	"fmt"
	"sync"
	"time"
)

func func1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("in func1")
}

func func2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("in func2")
	time.Sleep(2 * time.Second)
}
func printTable(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("number is", number)
}
func main() {
	var wg, wg1 sync.WaitGroup
	wg.Add(2)
	go func1(&wg)
	go func2(&wg)
	wg.Wait() //wait for both then

	for i := 1; i <= 10; i = i + 1 {
		wg1.Add(1)
		go printTable(i*2, &wg1)
	}
	wg1.Wait()
}
