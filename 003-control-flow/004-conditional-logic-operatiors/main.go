package main

import (
	"fmt"
)

func main() {
	// result : true
	fmt.Println(true && true)

	// result : false
	fmt.Println(true && false)

	// result : true
	fmt.Println(true || true)

	// result : true
	fmt.Println(true || false)

	// result : false
	fmt.Println(!true)
}
