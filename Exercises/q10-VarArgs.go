package main

import "fmt"

func main() {
	prt(1, 3, 4, 5, 6)
}

func prt(numbers ... int) {
	for _, v := range numbers {
		fmt.Printf("%d\n", v)
	}
}
