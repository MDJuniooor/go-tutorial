package main

import "fmt"

type hotdog5 int

var x5 hotdog5

func main() {
	x5 = 28
	y5 := int(x5)
	fmt.Println(y5)
	fmt.Printf("%T\n", y5)
}
