package main

import "fmt"

type calculate interface{
	getArea() float64
}

type shape interface{
	printArea() string
}

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

func printArea(c calculate) {
	fmt.Println(c.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	tri := triangle {}
	squ := square {}

	tri.base = 12.12
	tri.height = 22.22
	squ.sideLength = 10

	printArea(tri)
	printArea(squ)
}