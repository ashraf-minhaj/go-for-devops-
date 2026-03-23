package main

import "fmt"

// easy version
/*
func main() {
	run := true

	for run {
		fmt.Println("shit")
	}
}
*/

// the one that has a escape condition
func main() {
	run := true

	counter := 1

	for run {
		fmt.Println("shit")

		counter += 1

		if counter > 5 {
			run = false
			fmt.Println("counter is greater than 5, stopping the loop")
		}
	}
}
