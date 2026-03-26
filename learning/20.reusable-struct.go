package main

import "fmt"

// by declaring the struct with type, we made it reusable
type record struct {
	name    string
	age     int
	married bool
}

func main() {
	minhaj := record{
		name:    "ashraf minhaj",
		age:     900,
		married: true,
	}

	kabir := record{
		name:    "kabir ahmed",
		age:     44,
		married: false,
	}

	_ = kabir

	fmt.Println(minhaj)
	fmt.Println(minhaj.married)

	if minhaj.married {
		fmt.Println("Sorry ladies")
	}

}
