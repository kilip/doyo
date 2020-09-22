package main

import (
	"fmt"

	"rsc.io/quote"
)

// Hello function
func Hello() string {
	return "hello world"
}

func main() {
	quote.Hello()
	fmt.Println("Hello WOrld!")
}
