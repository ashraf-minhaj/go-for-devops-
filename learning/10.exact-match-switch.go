package main

func main() {
	a := 5

	switch a {
	case 1:
		println("a is 1")
	case 2:
		println("a is 2")
	case 3:
		println("a is 3")
	case 4:
		println("a is 4")
	case 5, 6: // can have multiple values for a case
		println("a is 5 or 6")
	default:
		println("a is something else")
	}
}
