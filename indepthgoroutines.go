package main

import (
	"fmt"
)

func main() {
	/*ch := make(chan int) //sending int

	//ch <- 1//main go rouine ending
	//send from other gorouin recieve from main
	go func() {
		ch <- 1
		time.Sleep(1 * time.Second)
	}()
	fmt.Println("sent succ")
	data := <-ch //recieving from main go rouine which block because sender needs reciver
	//problem solves wen we eiether send from other gorouine and main will recieve it or viceverase
	//time.Sleep(2 * time.Millisecond)

	//send from main goutine and recieve from other gorouitne
	//var data int = 0
	//go func() {
	//	data = <-ch
	//}()
	//ch <- 1

	fmt.Println("data is ", data)



		//Buffer
		//sendinf from main goruine and recieve from main goruoine

		ch := make(chan int, 3) // Buffered channel with capacity 1

		// Send a value into the channel
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("sent succ")

		// Receive the value from the channel
		data1 := <-ch
		data2 := <-ch
		data3 := <-ch
		//time.Sleep(2 * time.Millisecond)
		fmt.Println("data is", data1, data2, data3)
	*/

	go func() {
		fmt.Println("in go")
	}()
	fmt.Println("in main")
}
