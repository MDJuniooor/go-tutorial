package main

import "fmt"

func main4() {
	year := 1993

	for {
		if year == 2021 {
			break
		}
		fmt.Println(year)
		year++
	}
}
