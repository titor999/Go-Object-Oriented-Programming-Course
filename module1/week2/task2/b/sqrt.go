package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Input an integer >>>")
	var integer uint
	_, err := fmt.Scan(&integer)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Sqrt(integer))
}

func Sqrt(a uint) uint {
	var b = a
	var i = 0
	for b*b > a && i < 200 {
		b = (b + a/b) / 2
		i++
	}
	return b
}
