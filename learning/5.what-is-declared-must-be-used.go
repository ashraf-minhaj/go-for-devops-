package main

import (
	"fmt"
	_ "sync" // imported but not used
)

func main() {
	// variable declared but not used
	var a int = 10

	_ = a // assigned to nothing

	fmt.Println(a)
}
