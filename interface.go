package main

import "fmt"

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Put(v int) {
	p.i = v
}

func (p *S) Puta(v int) {
	p.i = v
}

type I interface {
	Get() int
	Put(i int)
}

func main() {
	//exe(p I)
}

func exe(i I) {
	fmt.Print(i.Get())
}
