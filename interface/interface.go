package main

import "fmt"

type Shape interface {
	area() int
	perim() int
}

type Rect struct {
	length, width int
}

type Square struct {
	side int
}

func (r Rect) area() int {
	return r.length * r.width
}

func (r Rect) perim() int {
	return (r.length + r.width) * 2
}

func (sq Square) area() int {
	return sq.side * sq.side
}

func (sq Square) perim() int {
	return sq.side * 2
}

func main() {
	var s Shape
	r := Rect{length: 5, width: 4}
	s = Shape(r)
	fmt.Println("Area: ", s.area())       // Area: 15
	fmt.Println("Perimeter: ", s.perim()) // Perimeter: 16

	q := Square{side: 5}
	s = q
	fmt.Println("Area: ", s.area())       // Area: 25
	fmt.Println("Perimeter: ", s.perim()) // Perimeter: 10
}
