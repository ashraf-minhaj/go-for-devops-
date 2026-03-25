package main

import "fmt"

func main() {
	fruits := [4]string{"apple", "banana", "orange", "grape"}

	fmt.Println(fruits[0])

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fruits[1] = "pineapple"

	// indx := []int[]
	// values := []string[]
	// indx, values := range(fruits)
	// fmt.Println(indx, values)

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	fruits[4] = "melon" // this will cause an error because array has fixed size of 4
}
