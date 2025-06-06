package main

import (
	"fmt"
	"time"
)

func sendWithUnbufferedRoutine1(ch chan int) {
	ch <- 1
}

func recieveWithUnbufferedRoutine1(ch chan int) {
	fmt.Println(<-ch)

}

func sendWithBufferedRoutine1(chBuffer chan int) {
	chBuffer <- 100
	chBuffer <- 200
}

func recieveWithBufferedRoutine1(ch chan int) {
	fmt.Println(<-ch)

}

func main() {

	ch := make(chan int)
	go sendWithUnbufferedRoutine1(ch)
	fmt.Println(<-ch)
	go recieveWithUnbufferedRoutine1(ch)

	ch <- 2

	go sendWithUnbufferedRoutine1(ch)
	go recieveWithUnbufferedRoutine1(ch)
	//waiting because sending from one go rotine to other routine
	time.Sleep(100 * time.Millisecond)

	chBuffer := make(chan int, 2)
	sendWithBufferedRoutine1(chBuffer) //works withot go

	fmt.Println(<-chBuffer, " ", <-chBuffer)

	//<-chBuffer//will trhow eror as there is no value in queue

	go recieveWithBufferedRoutine1(chBuffer) //listen
	go sendWithBufferedRoutine1(chBuffer)

	time.Sleep(100 * time.Millisecond)

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
	*/

	//Buffer
	//sendinf from main goruine and recieve from main goruoine

	/*ch := make(chan int, 3) // Buffered channel with capacity 1

	// Send a value into the channel
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("sent succ")

	// Receive the value from the channel
	data1 := <-ch
	data2 := <-ch
	data3 := <-ch //its ok if we dont recieve 3rd value
	//time.Sleep(2 * time.Millisecond)
	fmt.Println("data is", data1, data2, data3)

	ch1 := make(chan int, 1) // Buffered channel with capacity 2

	go func() {

		<-ch1
		<-ch1
		<-ch1
		<-ch1
	}()
	//if we send first before go will throe deadlock
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	ch <- 4*/

	//in below case we need time.sleep
	go func() {
		fmt.Println("in go")
	}()
	fmt.Println("in main")

}
