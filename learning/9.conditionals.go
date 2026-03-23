package main

import "fmt"

func someFunc() int {
	return 123
}

func main() {
	a := 5

	if a > 5 {
		fmt.Println("a is greater thatn 5")
	} else if a == 5 {
		fmt.Println("a is equal to 5")
	} else {
		fmt.Println("a is less than 5")
	}

	if value := someFunc(); value != 0 {
		fmt.Println("value is:", value)
	}
}
