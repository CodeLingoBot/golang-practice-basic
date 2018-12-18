package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}

func main() {
	var s stack
	s.push(25)
	fmt.Printf("stack %v\n", s)
	s.push(3)
	fmt.Printf("stack %v\n", s)
	s.pop()
	fmt.Printf("stack %v\n", s)
}

//func (s *stack) push(k int) {
//	if s.i+1 > 9 {
//		return
//	}
//	s.data[s.i] = k
//	s.i++
//}

func (s *stack) push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() {
	s.i--
	s.data[s.i] = 0
}
