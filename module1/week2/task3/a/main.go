package main

import "fmt"

func main() {
	fmt.Println("Input the array length >>>")
	var arrayLength int
	_, _ = fmt.Scan(&arrayLength)

	var names = make([]int, arrayLength)
	fmt.Println("Input numbers >>>")
	var arr int

	for i := 0; i < arrayLength; i++ {
		_, _ = fmt.Scan(&arr)
		names[i] = arr
	}

	max := maxValue(names)

	fmt.Println(max)
}

func maxValue(arr []int) int {
	var max = arr[0]
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
