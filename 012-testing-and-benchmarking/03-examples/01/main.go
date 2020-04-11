package main

import (
	"fmt"
	"github.com/GoesToEleven/forexample/012-testing-and-benchmarking/03-examples/01/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
