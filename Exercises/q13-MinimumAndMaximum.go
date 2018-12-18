package main

import "fmt"

func main() {
	arr := []int{6, 3, 7, 1, 10}
	max := max(arr)
	min := min(arr)
	fmt.Println(max)
	fmt.Println(min)
}

func max(numbers []int) (max int) {
	max = numbers[0]
	for _, v := range numbers {
		if v > max {
			max = v
		}
	}
	return max
}

func min(numbers []int) (min int) {
	min = numbers[0]
	for _, v := range numbers {
		if v < min {
			min = v
		}
	}
	return min
}
