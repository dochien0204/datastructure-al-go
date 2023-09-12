package main

import (
	"fmt"
)

func Display(arr []int) {
	for _, val := range arr {
		fmt.Print(fmt.Sprintf("%d ", val))
	}
	fmt.Println()
}

func FillChar(arr []int, index int) {
	arr[index] = 1
	for i := len(arr) - 1; i > index; i-- {
		arr[i] = 0
	}
}

func main() {
	n := 3
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 0
	}

	arrRef := &arr
	for {
		Display(*arrRef)
		i := n - 1
		for i >= 0 && arr[i] > 0 {
			i--
		}

		if i >= 0 {
			arr[i] = 1
			FillChar(arr, i)
		}

		if i < 0 {
			break
		}
	}
}
