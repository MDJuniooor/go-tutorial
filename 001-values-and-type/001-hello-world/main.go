package main

import "fmt"

func main() {

	fmt.Println("hello world")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	foo()
	bar()
}

func foo() {
	fmt.Println("I'm a foo")
}

func bar() {
	fmt.Println("I'm a bar")
}
