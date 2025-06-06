package main

import (
	"fmt"
	"os"
)

func test(arr []int) {
	defer fmt.Print("last")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()
	index := 2
	if index > 1 {
		//	panic("out of bound")
	}
	fmt.Println(arr[2])
	fmt.Println("will work")
}

func readFile() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed when the function exits

	fmt.Println("File opened successfully")

}
func main() {
	arr := []int{1, 2}
	test(arr)

	readFile()
	fmt.Println("in main")
}
