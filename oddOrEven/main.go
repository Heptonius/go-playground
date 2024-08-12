package main

import "fmt"

const ODD = "odd"
const EVEN = "even"

func oddOrEven(number int) string {
	if number%2 == 1 {
		return ODD
	} else {
		return EVEN
	}
}

func main() {
	numbers := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	}

	for _, number := range numbers {
		fmt.Println(oddOrEven(number))
	}
}
