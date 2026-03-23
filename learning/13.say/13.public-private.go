package main

import (
	"fmt"
	"learning/say"
)

func main() {
	say.PrintHello()
	// say.printWorld() // says undefined - because it's private

	fmt.Println(say.OutsideValue)

	// fmt.Println(say.OutsideValue2) // says undefined - because it's private
	// fmt.Print(say.PrintHello().InsideValue) // not possible because InsideValue is not accessible outside of the function
}
