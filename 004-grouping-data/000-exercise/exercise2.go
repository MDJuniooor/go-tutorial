package main

import "fmt"

func main2() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range x[3:] {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
	fmt.Println(cap(x))
}
