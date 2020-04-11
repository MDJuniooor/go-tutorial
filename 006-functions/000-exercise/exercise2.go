package main

import "fmt"

func main2() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo2(ii...)
	fmt.Println(n)

	n2 := bar2(ii)
	fmt.Println(n2)
}

func foo2(xi ...int) int {
	yi := 0
	for _, vi := range xi {
		yi += vi
	}
	return yi
}

func bar2(xi []int) int {
	yi := 0
	for _, vi := range xi {
		yi += vi
	}
	return yi
}
