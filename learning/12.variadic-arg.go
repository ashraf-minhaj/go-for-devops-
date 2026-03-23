package main

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	println(sum([]int{1, 2, 3, 4, 5}))
}
