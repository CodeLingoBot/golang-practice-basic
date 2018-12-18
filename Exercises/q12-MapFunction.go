package main

import "fmt"

func main() {
	m := []int{1, 2, 3, 4, 5}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", Map(f, m))
}

func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}
