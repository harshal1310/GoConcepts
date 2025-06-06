package main

import (
	"fmt"
)

type Box[T any] struct {
	Value T
	shape T
}

func (b Box[T]) Print() {
	fmt.Println("Box contains:", b.Value)
}

func main() {
	intBox := Box[int]{Value: 42}
	strBox := Box[string]{Value: "Fiber"}

	intBox.Print()
	strBox.Print()
}
