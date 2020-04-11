package main

import (
	"fmt"
)

var x int
var y float64

// int8 range -128 ~ 127
var z int8 = 127

// byte eg.. utf-8
var a byte

// rune eg.. unicode
var b rune

func main() {
	x = 42
	y = 42.34534
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}
