package main

import "fmt"

func main4() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}

		fmt.Println(x)
	}
	fmt.Println("done.")
}
