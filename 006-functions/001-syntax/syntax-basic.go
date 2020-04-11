package main

import (
	"fmt"
)

func main1() {
	foo()
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("hello from foo")
}
