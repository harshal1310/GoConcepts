package main

import (
	"fmt"
	"reflect"
	"time"
)

func reflectExample() {
	x := 42
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println("Type:", t)
	fmt.Println("Value:", v)
}

func main() {
	reflectExample()

	ch := make(chan int)

	go send(ch)

	time.Sleep(2 * time.Millisecond)

	//for i := 0; i < 5; i++ {

	//	num := <-ch
	//	fmt.Println("receiving", num)
	//}
	for i := range ch {
		fmt.Println("receiving", i)
	}

}

func send(ch chan int) {
	fmt.Println("sending")
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("sending", i)
	}
	close(ch)
}
