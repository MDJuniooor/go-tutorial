package main

import (
	"fmt"
)

func main4() {
	x := bar()

	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)

}

func bar() func() int {
	return func() int {
		return 451
	}
}
