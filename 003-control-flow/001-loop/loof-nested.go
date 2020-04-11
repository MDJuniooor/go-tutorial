package main

import "fmt"

func main2() {
	for i := 0; i < 10; i++ {
		fmt.Println("outer loop")
		for j := 0; j < 10; j++ {
			fmt.Println("inner loop")
			fmt.Println(i, j)
		}
	}
}
