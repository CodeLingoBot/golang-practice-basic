package main

import "fmt"

type Vuong struct {
	length, width int
}

type VuongWithHeight struct {
	Vuong
	height int
}

func (r Vuong) area() int {
	return r.length * r.width
}

func (r Vuong) perim() int {
	return (r.length + r.width) * 2
}

func (r VuongWithHeight) sum() int{
	return r.length + r.width + r.height
}

func main() {
	ra := VuongWithHeight{Vuong{2,3}, 5}
	fmt.Printf("area: %d \n", ra.area())
	fmt.Printf("perim: %d \n", ra.perim())
	fmt.Printf("perim: %d \n", ra.sum())
}
