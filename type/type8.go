package main

import "fmt"

type PInterface interface{
	TestParent1(i int) int
}

type Parent struct {
	name string
}

type Child struct {
	Parent
	age int
}

func TestChild(c Child){
	fmt.Println(c.age)
}

func TestParent(p Parent){
	fmt.Println(p.name)
}

func (p Parent) TestParent1(i int) int{
	fmt.Println(i)
	return i+1
}


func testInterface(pInterface PInterface){

}

func main(){
	p := Parent{name: "parent"}
	//Test(p) //cannot use p (type Parent) as type Child in argument to Test
	c :=Child{age:18, Parent: Parent{name:"parent"}}
	TestChild(c)
	TestParent(c.Parent)

	testInterface(c)
	testInterface(p)

	t:= p.TestParent1 // method value
	fmt.Println(t(9))

	y := Parent.TestParent1 // method expression
	fmt.Println(y(p, 100))
}