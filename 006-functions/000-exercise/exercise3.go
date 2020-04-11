package main

import "fmt"

func main3() {
	defer foo3()
	fmt.Println("hello playground")
}

func foo3() {
	defer func() {
		fmt.Println("foo defer ran")
	}()
	fmt.Println("foo ran")
}
