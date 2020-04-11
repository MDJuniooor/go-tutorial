package main

import "fmt"

func main6() {
	sum := func(a int, b int) int {
		return a + b
	}
	fmt.Println(sum(2, 6))
}
