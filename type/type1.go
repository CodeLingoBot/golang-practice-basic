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
	fmt.Println("Vuong area")
	return r.length * r.width
}

func (r Vuong) perim() int {
	fmt.Println("Vuong perim")
	return (r.length + r.width) * 2
}

func (r VuongWithHeight) sum() int{
	fmt.Println("VuongWithHeight sum")
	return r.length + r.width + r.height
}

func (r VuongWithHeight) area() int{
	fmt.Println("VuongWithHeight area")
	return r.length + r.width + r.height
}

func main() {
	ra := VuongWithHeight{Vuong{2,3}, 5}
	fmt.Printf("area: %d \n", ra.area())
	fmt.Printf("perim: %d \n", ra.perim())
	fmt.Printf("perim: %d \n", ra.sum())
}
