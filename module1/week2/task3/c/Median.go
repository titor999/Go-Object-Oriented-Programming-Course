package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Input the array length >>>")
	var arrayLength int
	_, _ = fmt.Scan(&arrayLength)

	var arr = make([]int, arrayLength)
	fmt.Println("Input elements array >>>")
	var element int

	for i := 0; i < arrayLength; i++ {
		_, _ = fmt.Scan(&element)
		arr[i] = element
	}

	sort.Ints(arr)
	fmt.Println(arr)

	fmt.Println(median(arr))

}

func median(arr []int) int {
	var median int
	if len(arr)%2 == 0 {
		median = (arr[len(arr)/2] + arr[len(arr)/2-1]) / 2
	} else {
		median = arr[len(arr)/2]
	}
	return median
}
