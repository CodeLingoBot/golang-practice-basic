package main

import "fmt"

type asdad struct {
	name string
	age int
}

type response struct {
	fname string
	fage int
}

func main(){
	a := asdad{name:"h", age:1}
	b:=format(a)
	fmt.Println(b)
}

func format(a asdad)(*response){
	res := new(response)
	res.fname = a.name +"iep"
	res.fage = a.age + 1
	return res
}