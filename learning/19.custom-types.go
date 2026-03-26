package main

import (
	"fmt"
	"reflect"
)

func main() {
	// create a type
	type CarModel string

	// cerate a variable with that type
	var car CarModel = "audi"

	fmt.Println(reflect.TypeOf(car))

}
