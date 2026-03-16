package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(10, 20)
	fmt.Println(sum)
}
