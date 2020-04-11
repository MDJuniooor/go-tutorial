package main

import (
	"fmt"
)

func main2() {
	foo()
	bar("James")
}

func foo() {
	fmt.Println("hello from foo")
}

// everything in Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,", s)
}
