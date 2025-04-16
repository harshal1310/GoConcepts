package main

import (
	"fmt"
	"reflect"
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
}
