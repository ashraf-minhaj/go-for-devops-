package main

import "fmt"

func main() {
	var record = struct {
		name string
		age  int
	}{
		name: "minhaj",
		age:  900,
	}

	fmt.Println(record.name)
	fmt.Println(record.age)
}
