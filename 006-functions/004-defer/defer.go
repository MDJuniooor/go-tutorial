package main

import (
	"fmt"
)

func main() {
	// i don't know exactly defer conception...
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
