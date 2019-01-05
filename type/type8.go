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
    //Method values are useful when packages API calls for a function value,
    // and the client’s desired behavior for that function is to call a method on a specific receiver.
	t:= p.TestParent1 // method value
	fmt.Println(t(9))

	//Method expressions can be helpful when you need a value to represent
	// a choice among several methods belonging to the same type
	// so that you can call the chosen method with many different receivers
	y := Parent.TestParent1 // method expression
	fmt.Println(y(p, 100))

	u := (*Parent).TestParent1
	fmt.Println(u(&p, 1000))
}