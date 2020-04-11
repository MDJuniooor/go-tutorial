package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func main4() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p1.speak()
}
