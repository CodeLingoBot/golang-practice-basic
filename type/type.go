package main

import "fmt"

type Rect struct {
	length, width int
}

func (r Rect) area() int {
	return r.length * r.width
}

func (r Rect) perim() int {
	return (r.length + r.width) * 2
}

func main() {
	r := Rect{length: 4, width: 5}
	fmt.Printf("area: %d \n", r.area())
	fmt.Printf("perim: %d", r.perim())
}
