package main

import "fmt"

func main() {
	count(1)
	//fmt.Println("a")
}

func count(i int) {
	if i == 10 {
		return
	}
	count(i+1)
	fmt.Printf("%d ", i)
}
