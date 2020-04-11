package main

import "fmt"

var x3 int = 4
var y3 string = "james bond"
var z3 bool = true

func main3() {
	s := fmt.Sprintf("%v\t%v\t%v", x3, y3, z3)
	fmt.Println(s)
}
