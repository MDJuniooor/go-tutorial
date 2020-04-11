package main

import "fmt"

func main3() {
	x := 0
	for x < 10 {
		if x == 7 {
			break
		}
		fmt.Println(x)
		x++
	}
}
