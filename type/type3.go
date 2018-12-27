package main

import "fmt"

//var x int
func main(){
	x := 1993
	p := &x
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(&p)
	fmt.Println(&x)
	fmt.Println(*p)
	//fmt.Println(*x) invalid indirect of x (type int)
	fmt.Println(p == &x)
	fmt.Println(*p == x)
}