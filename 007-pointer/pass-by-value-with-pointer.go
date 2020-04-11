package main

import (
	"fmt"
)

func main6() {
	x := 0
	fmt.Println("x befor", &x)
	fmt.Println("x befor", x)
	foo2(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo2(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
