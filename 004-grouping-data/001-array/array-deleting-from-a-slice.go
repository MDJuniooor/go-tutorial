package main

import "fmt"

func main6() {
	x := []int{4, 5, 6, 7, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1024)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
