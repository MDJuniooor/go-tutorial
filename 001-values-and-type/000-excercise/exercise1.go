package main

import "fmt"

var x int
var y string
var z bool

func main1() {
	x = 42
	y = "james bond"
	z = true
	fmt.Println("single statement")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("multiple statement")
	fmt.Println(x, y, z)
}
