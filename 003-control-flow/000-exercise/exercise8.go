package main

import (
	"fmt"
)

func main8() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}
