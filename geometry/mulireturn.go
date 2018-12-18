package main

import "fmt"

func main(){
	a , b, c := multireturn(4, 3, 2)
	fmt.Print(a, b, c)
}

func multireturn(i int, i2 , i3 int) (int , int, int) {
	return i+1, i2+2, i3+3
}

