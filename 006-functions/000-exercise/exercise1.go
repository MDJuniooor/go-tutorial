package main

import "fmt"

func main1() {
	x := foo1()
	y1, y2 := bar1()
	fmt.Println(x)
	fmt.Println(y1, y2)
}

func foo1() int {
	return 100
}

func bar1() (string, int) {
	return "test", 1
}
