package main

import "fmt"

func main7() {
	f := func(a int) int {
		return a
	}
	fmt.Println(f(1))
	g := f
	fmt.Println(g(2))
}
