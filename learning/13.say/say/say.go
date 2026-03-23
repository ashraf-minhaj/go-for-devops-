package say

import "fmt"

var OutsideValue int = 10
var outsideValue2 int = 30

func PrintHello() {

	InsideValue := 20

	fmt.Println("Hello", InsideValue)
}

func printWorld() {
	fmt.Println("World")
}
