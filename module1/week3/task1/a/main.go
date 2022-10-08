package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("What line are we looking for?")
	var searchLine string
	_, _ = fmt.Scan(&searchLine)

	fmt.Println("How many lines do I get?")
	var linesCount int
	_, _ = fmt.Scan(&linesCount)

	var rows string
	fmt.Println("Enter the rows in which you want to search")
	var row string

	for i := 0; i < linesCount; i++ {
		_, _ = fmt.Scan(&row)
		rows += row
	}

	result := strings.Count(rows, searchLine)
	fmt.Println(result)
}
