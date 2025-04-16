package main

import (
	"fmt"
)

func test() {
	defer fmt.Print("last")

	fmt.Println("first")
}
func main() {
	test()
}
