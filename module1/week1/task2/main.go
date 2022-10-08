package main

import (
	"fmt"
	"math"
)

func main() {
	// const uint for all types (uint8, uint16, uint32, uint64, uint)
	const (
		MinUint      = 0
		MinByte byte = 0
		MaxByte byte = 255
		MinRune rune = math.MinInt32
		MaxRune rune = math.MaxInt32
	)

	fmt.Println("Type\t\tMin Size\t\t\tMax Size")

	fmt.Printf("%[1]T\t\t%[1]d\t\t\t\t%[1]d\n", int8(math.MinInt8), int8(math.MaxInt8))
	fmt.Printf("int16\t\t%d\t\t\t\t%d\n", math.MinInt16, math.MaxInt16)
	fmt.Printf("int32\t\t%d\t\t\t%d\n", math.MinInt32, math.MaxInt32)
	fmt.Printf("int64\t\t%d\t\t%d\n", math.MinInt64, math.MaxInt64)
	fmt.Printf("int\t\t%d\t\t%d\n\n", math.MinInt, math.MaxInt)

	fmt.Printf("uint8\t\t%d\t\t\t\t%d\n", MinUint, math.MaxUint8)
	fmt.Printf("uint16\t\t%d\t\t\t\t%d\n", MinUint, math.MaxUint16)
	fmt.Printf("uint32\t\t%d\t\t\t\t%d\n", MinUint, math.MaxUint32)

	/*
		Without an explicit type,
		by default, integer untyped constants will get
		an int type, which can't hold math.MaxUint64.
		We need to use "uint64(math.MaxUint64)" and "uint(math.MaxUint)"
	*/
	fmt.Printf("uint64\t\t%d\t\t\t\t%d\n", MinUint, uint64(math.MaxUint64))
	fmt.Printf("uint\t\t%d\t\t\t\t%d\n\n", MinUint, uint(math.MaxUint))

	fmt.Printf("float32\t\t%[1]v\t\t%[2]v\n", math.SmallestNonzeroFloat32, math.MaxFloat32)
	fmt.Printf("float64\t\t%[1]v\t\t\t\t%[2]v\n\n", math.SmallestNonzeroFloat64, math.MaxFloat64)

	fmt.Printf("%[1]T\t%[1]v\t\t\t%[1]v\n",
		complex(float32(math.SmallestNonzeroFloat32), float32(math.SmallestNonzeroFloat32)),
		complex(float32(math.MaxFloat32), float32(math.MaxFloat32)),
	)

	fmt.Printf("%[1]T\t%[1]v\t\t%[1]v\n\n",
		complex(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64),
		complex(math.MaxFloat64, math.MaxFloat64),
	)

	fmt.Printf("byte\t\t%[1]d\t\t\t\t%[1]d\n", MinByte, MaxByte)
	fmt.Printf("rune\t\t%[1]d\t\t\t%[1]d\n", MinRune, MaxRune)

}
