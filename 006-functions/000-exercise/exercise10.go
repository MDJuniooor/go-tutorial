package main

import "fmt"

func main() {
	p1 := get()
	p2 := get()
	fmt.Println(p1())
	fmt.Println(p1())
	fmt.Println(p1())
	fmt.Println(p2())
	fmt.Println(p2())
}

func get() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
