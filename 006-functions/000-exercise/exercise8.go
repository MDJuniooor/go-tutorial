package main

import "fmt"

func main8() {
	ii := []int{1, 2, 3, 4}
	f1 := highOrderFunc(func(a ...int) int {
		return len(a)
	}, 1)
	fmt.Println(f1(ii...))
}

func highOrderFunc(f func(a ...int) int, value int) func(a ...int) int {
	fmt.Println("high Order Func : ", value)
	return f
}
