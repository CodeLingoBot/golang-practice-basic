package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1, 4, 6}
	bubblesort(arr)
	fmt.Println(arr)
}

func bubblesort(arr [] int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}
