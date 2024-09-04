// package main

// import "fmt"

// type triangle struct {
// 	height float64
// 	base   float64
// }

// type square struct {
// 	sideLength float64
// }

// type shape interface {
// 	getArea() float64
// }

// func (t triangle) getArea() float64 {
// 	return t.base * t.height / 2
// }

// func (s square) getArea() float64 {
// 	return s.sideLength * s.sideLength
// }

// func printArea(s shape) {
// 	fmt.Println(s.getArea())
// }

// func main() {
// 	tri := triangle{
// 		height: 15,
// 		base:   20,
// 	}
// 	sqr := square{
// 		sideLength: 10,
// 	}

// 	printArea(tri)
// 	printArea(sqr)
// }
