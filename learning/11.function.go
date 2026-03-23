package main

func sum(a int, b int) int {
	return a + b
}

func multiple(a int, s string) (result int, name string) {
	result = a * 2
	name = s + " boss"

	return result, name
}

func main() {
	println(sum(2, 3))

	res, na := multiple(2, "minhaj")
	println(res, na)
}
