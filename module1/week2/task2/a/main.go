package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Input the pyramid height >>>")
	var pyramidHeight int
	_, _ = fmt.Scan(&pyramidHeight)

	for i := 0; i <= pyramidHeight; i++ {
		// Left Pattern
		for j := i; j < pyramidHeight; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("#")
		}

		// Space between left and right pyramids
		fmt.Print(" ")

		// Right Pattern
		fmt.Println(strings.Repeat("#", i))
	}
}
