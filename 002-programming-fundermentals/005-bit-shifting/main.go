package main

import (
	"fmt"
)

func main1() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}

func main2() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
