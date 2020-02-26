package main

import "fmt"

type shape interface {
	getArea() float64
}

// make printAra a function usuable to all structs that implement the shape interface
func printArea(s shape) {
	fmt.Println(s.getArea())
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	length float64
}

func (t triangle) getArea() float64 {
	return ((t.base * t.height) / 2)
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func main() {
	myTriangle := triangle{
		base:   4,
		height: 5,
	}

	mySquare := square{length: 20}

	printArea(myTriangle)
	printArea(mySquare)

	fmt.Println("using tringle's getMethod:", myTriangle.getArea())
	fmt.Println("using square's getMethod:", mySquare.getArea())

}
