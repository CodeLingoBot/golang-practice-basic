package main

import "fmt"

func sum(vals ... int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func a() int{
	i := 0
	defer fmt.Println(i)
	i++
	return i
}

func f(){
	fmt.Println("a")
}

func main() {
	defer fmt.Println("c")
	defer f()
	fmt.Println("b")

}
