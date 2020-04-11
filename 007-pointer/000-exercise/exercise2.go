package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "changed first"
}

func main() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   20,
	}
	fmt.Println(p1.first)
	changeMe(&p1)
	fmt.Println(p1.first)
}
