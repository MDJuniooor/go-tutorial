package main

import "fmt"

func main3() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a

	fmt.Println(b)
	fmt.Println(*b)
	// * is dereference operator
}
