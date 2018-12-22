package main

import "fmt"

type People struct {
	name string
	age  int
}

func main() {
	hiep := new(People)
	hiep.name = "Hiep";
	hiep.age = 18
	fmt.Printf("%v\n", hiep)

	mystring := "this is a string"
	byteslice := []byte(mystring)
	fmt.Print(byteslice)

	slice := []rune(mystring)

	fmt.Print(slice)
}
