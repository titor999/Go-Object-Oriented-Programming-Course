package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Input an Icosahedron volume >>>")
	var ribLength float64
	_, err := fmt.Scan(&ribLength)
	if err != nil {
		return
	}
	fmt.Println(GetVolumeIcosahedron(ribLength))
}

func GetVolumeIcosahedron(ribLength float64) float64 {
	var volumeIcosahedron = 5 * math.Pow(ribLength, 3) / 12 * (3 + math.Sqrt(5))
	return volumeIcosahedron
}
