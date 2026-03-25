package main

import (
	"fmt"
	"reflect"
)

func main() {
	namePerRoll := map[int]string{}

	// insert values
	namePerRoll[1] = "minhaj"
	namePerRoll[33] = "ashraf"
	namePerRoll[55] = "sami"
	namePerRoll[77] = "ali"
	namePerRoll[99] = "hassan"
	namePerRoll[100] = "hussain"
	namePerRoll[101] = "mohammad"
	namePerRoll[102] = "ahmed"

	fmt.Println(namePerRoll)
	fmt.Println(namePerRoll[33])

	// get length of map
	numOfValues := len(namePerRoll)
	fmt.Println(numOfValues)

	// traverse
	for roll, name := range namePerRoll {
		fmt.Println("Person", name, "has roll", roll)

		if name == "minhaj" {
			fmt.Println("Hell yeah, boss", name, "is found at roll", roll)
			fmt.Println("%T\n", name)         // does not work
			fmt.Println(reflect.TypeOf(name)) // works
		}
	}

	// or use ok to find a value
	name, ok := namePerRoll[33]
	if ok {
		// Key exists, use 'value'
		fmt.Println("Found name:", name)
	} else {
		// Key does not exist
		fmt.Println("Name not found")
	}

}
