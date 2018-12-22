package main

import (
	"./even"
	"./stack"
	"fmt"
)

func main() {
	i := 5
	fmt.Printf("Is %d even? %v\n", i, even.Even(i))
	st := stack.Stack{}
	st.Push(2);
	fmt.Println(st)
	st.PopByIndex(0)
	fmt.Println(st)
	st.Push(4)
	fmt.Println(st)
	st.Pop()
	fmt.Println(st)
}
