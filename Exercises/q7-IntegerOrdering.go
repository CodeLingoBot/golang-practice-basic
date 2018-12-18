package main

import "fmt"

func main() {
	first, second := order(100, 121)

	fmt.Println(first, second)
}

func order(i1, i2 int) (int, int) {
	if i1 > i2 {
		return i2, i1
	} else {
		return i1, i2
	}
}
