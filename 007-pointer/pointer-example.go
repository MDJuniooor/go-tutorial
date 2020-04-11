package main

import "fmt"

func main2() {
	a := 28

	fmt.Println(a)
	fmt.Println(&a)

	var b = &a // pointer

	fmt.Println(b)
	// type is *int
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &a)

}
