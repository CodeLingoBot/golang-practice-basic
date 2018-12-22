package main

import "fmt"

type e interface {
}

func multi(f e) e {
	switch f.(type) {
	case int:
		return f.(int) * 2
	case string:
		return f.(string) + f.(string) + f.(string)
	}
	return f
}

func Mapping(f func(i e) e, arr []e) []e {
	m := make([]e, len(arr))
	for index, value := range arr {
		m[index] = f(value)
	}
	return m
}

func main() {
	m := []e{1, 2, 3, 4}
	mf := Mapping(multi, m)
	fmt.Printf("%v", mf)

	s := []e{"s", "t", "i", "n", "g"}
	ms := Mapping(multi, s)
	fmt.Printf("%v", ms)
}
