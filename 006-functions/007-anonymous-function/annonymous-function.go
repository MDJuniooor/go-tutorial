package main

import "fmt"

func main() {
	foo()
	func(x int) {
		fmt.Println("the meaning of life : ", x)
	}(28)
}

func foo() {
	fmt.Println("foo ran")
}
