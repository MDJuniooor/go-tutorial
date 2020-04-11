package main

import "fmt"

type hotdog int

var x4 hotdog

func main4() {
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Println(x4)
	fmt.Printf("%T\n", x4)

}
