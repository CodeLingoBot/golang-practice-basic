package main

import "fmt"
func main() {
	lst := []float64{1, 2, 3, 4, 5}
	avge := avg(lst)
	fmt.Println(avge)
}

func avg(lst []float64) (avg float64) {
	sum := 0.0
	if len(lst) > 0 {
		for _, v := range lst {
			sum += v
		}
		avg = sum / float64(len(lst))
	}
	return
}
