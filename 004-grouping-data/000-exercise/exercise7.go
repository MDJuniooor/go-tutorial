package main

import "fmt"

func main7() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	z := [][]string{x, y}
	fmt.Println(z)

	for i, v := range z {
		fmt.Println(i, v)
		for i2, v2 := range v {
			fmt.Println(i2, v2)
		}
	}
}
