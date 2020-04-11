package main

import "fmt"

// struct is a composite data type aka aggregate (complex) data types

type person struct {
	first string
	last  string
	age   int
}

func main1() {
	p1 := person{
		first: "james",
		last:  "bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
