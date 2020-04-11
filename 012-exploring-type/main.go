package main

import "fmt"

// implicit declaration
var y = 42

var z string = "shaken, not stireed"

var a string = `James said,
"shaken, 

not stireed"`

// this is a STATIC programming language
// a VARIABLE IS DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // Type print
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(y)
	fmt.Printf("%T\n", a)

}
