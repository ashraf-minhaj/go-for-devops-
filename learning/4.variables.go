package main

import "fmt"

// variables declared
// package scoped
var apple int = 10

// group of variables
// package scoped
var (
	banana int // value is 0 by default
	word   = "potato"
)

// create and assign
// should not be possible at package level
// potato := "pohtatoh"

func main() {
	fmt.Println(apple)
	fmt.Println(banana)
	fmt.Println(word)

	// function scoped variable
	// create and assign
	potato := "pohtatoh"
	fmt.Println(potato)

	// statement scoped variable
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
}
