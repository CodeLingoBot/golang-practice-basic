package main

import "fmt"

//var x int
var z *int
func main(){
	x := 1993
	p := &x
	fmt.Println(x) // 1993
	fmt.Println(p) // 0xc000014050
	fmt.Println(&p) // 0xc00000c028
	fmt.Println(&x) // 0xc000014050
	fmt.Println(*p) // 1993
	//fmt.Println(*x) invalid indirect of x (type int)
	fmt.Println(p == &x) //true
	fmt.Println(*p == x) // true
	//z = &p cannot use &p (type **int) as type *int in assignment
	z = &x // 0xc000014050
	fmt.Println(z)

	var t int
	fmt.Println(t) //0
	fmt.Println(&t) //0xc000014080
}