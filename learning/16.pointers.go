package main

func main() {
	a := 5
	print(&a) // prints the memory address of the variable 'a'

	// now let's store the mem addr to a pointer variable

	// pointerToA := &a
	var pointerToA *int = &a

	println(pointerToA)  // prints the memory address of 'a'
	println(*pointerToA) // dereferencing the pointer to get the value of 'a'

	// now change value at the memory address instead of cahnging x
	*pointerToA = 10
	println(a) // prints 10, because we changed the value at the memory address of 'a' through the pointer
}
