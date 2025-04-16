package custom

import (
	"fmt"
)

func main() {
	fmt.Println("Order Service Running...")
	var age int = 25

	if age >= 18 {
		fmt.Println("You are eligible to vote", age)
	} else {
		fmt.Println("You are not eligible to vote", age)
	}
	const name = "John"

	floatvalue := 3.13
	fmt.Println(floatvalue)

	var employess string = "John, Peter, Jack"
	fmt.Println(employess)

	for i := 0.01; i < floatvalue; i += 1 {
		fmt.Println(i)
	}

	//while like loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	numbers := []int{1, 2, 3, 4, 5}
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	result := add(10, 5)
	fmt.Println("Sum:", result)

	a, b := mulipleRetrunValue()
	fmt.Println(a, b, 9)

	result = sum(1, 2, 3)

	fmt.Println("Sum:", result)

}
func add(a int, b int) int { // last int for return type
	return a + b
}

func mulipleRetrunValue() (int, int) {
	return 5, 10
}

func sum(numbers ...int) int {
	total := 0
	//_ laternative of index if not usin index
	for _, num := range numbers {
		total += num
	}
	return total
}
