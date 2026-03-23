package main

import "fmt"

func main() {
	// for loop with initialization, condition, and post statement
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	// fmt.Println("the final value of i is", i, "bro") // undefined error, outside of the scope

	fmt.Println("/n")
	// without init statement
	var j int
	for ; j < 5; j++ {
		fmt.Print(j)
	}

	// we can access j because it's declared outside of the loop scope
	// now it's a function scoped variable
	fmt.Println("the final value of j is", j, "bro")
}
