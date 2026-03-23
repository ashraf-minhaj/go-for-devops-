package main

// range returns index and value
// we are not using index, so we use _ to ignore it
/*
for _, number := range numbers {
		sum += number
	}
*/

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	println(sum(1, 2, 3, 4, 5))
}
