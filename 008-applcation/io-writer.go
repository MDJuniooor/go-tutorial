package main

import (
	"fmt"
	"io"
	"os"
)

func main3() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground")
	io.WriteString(os.Stdout, "Hello, playground")
}
