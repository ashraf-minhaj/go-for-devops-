package main

import "fmt"

// simple version
/*
func main() {
	i := 1

	for {
		fmt.Println(i)
		i++
	}
}
*/

// having escape condition
func main() {
	i := 1

	for {
		fmt.Println(i)
		i++

		if i > 5 {
			break
		}
	}
}
