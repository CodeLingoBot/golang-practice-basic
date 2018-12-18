package main

import "fmt"

func main(){
	callback(4, printit)
}

func printit(x int) {
	fmt.Println(2)
	fmt.Printf("%v\n", x)
}

func callback(y int, f func(int)) {
	fmt.Println(1)
	f(y)
}