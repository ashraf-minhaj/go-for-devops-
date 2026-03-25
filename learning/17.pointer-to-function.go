package main

func changeValue(a *int, value int) {
	*a = value
}

func main() {
	a := 5
	changeValue(&a, 10)
	println(a)
}
