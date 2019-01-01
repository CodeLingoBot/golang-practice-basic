package main

import "fmt"

type newyear struct {
	year int
}

func (y *newyear) print1(){
	y.year += 1
}

func (y newyear) print2(){
	y.year += 1
}

func main(){
	y := newyear{year:2018}
	y.print1()
	//y.print2()
	fmt.Println(y.year)
	y2 := &newyear{year:2018}
	//y2.print1()
	y2.print2()
	fmt.Println(*(y2))

	newyear{year:2018}.print2()
}