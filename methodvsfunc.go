package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    greet("GoLang") // Output: Hello, GoLang
}




//method
package main

import "fmt"

type Person struct {
    name string
}

func (p Person) greet() {
    fmt.Println("Hello,", p.name)
}

func main() {
    p := Person{name: "Harshal"}
    p.greet() // Output: Hello, Harshal
}
