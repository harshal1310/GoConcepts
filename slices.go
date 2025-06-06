// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4]  // includes index 1, 2, 3
	fmt.Println(slice) // [20 30 40]
	fmt.Println(arr)

	s := make([]int, 5)               // len=5, cap=5
	s2 := make([]int, 3, 10)          // len=3, cap=10
	fmt.Println(s, len(s), cap(s))    // [0 0 0 0 0] 5 5
	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0] 3 10

	base := []int{1, 2, 3, 4, 5}
	sub := base[1:4]  // [2 3 4]
	sub2 := sub[1:2]  // [3]
	fmt.Println(sub2) // valid as long as capacity allows

	s = []int{1, 2}
	s = append(s, 3, 4)
	fmt.Println(s) // [1 2 3 4]

	// change original slice
	arr1 := [3]int{1, 2, 3}
	s = arr1[0:3] //. s=arr1[:]
	fmt.Println(s)
	s[0] = 100
	fmt.Println(arr1) // [100 2 3]

	a := []int{1, 2, 3}
	b := a[:2]
	fmt.Println(b)
	b = append(b, 100, 200) // capacity exceeded
	b[0] = 999

	fmt.Println("a:", a) // unchanged
	fmt.Println("b:", b) // new array

	s1 := []int{1, 2, 3, 4, 5}
	i := 2 // remove index 2
	s1 = append(s1[:i], s1[i+1:]...)
	fmt.Println(s1) // [1 2 4 5]

	s3 := []string{"a", "b", "c"}
	for i, val := range s3 {
		fmt.Printf("Index %d = %s\n", i, val)
	}

	for i := 0; i < len(s3); i++ {
		fmt.Printf("Index %d = %s\n", i, s3[i])
	}

}
