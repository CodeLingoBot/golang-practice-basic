package main

type Areaer interface {
	area() int
}

type Perimer interface {
	perim() int
}

type Shaper interface {
	Areaer
	Perimer
}


