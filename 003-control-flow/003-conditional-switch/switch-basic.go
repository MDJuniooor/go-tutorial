package main

import (
	"fmt"
)

func main1() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (2 == 3):
		fmt.Println("prints")
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	default:
		fmt.Println("this is default")
	}
}
