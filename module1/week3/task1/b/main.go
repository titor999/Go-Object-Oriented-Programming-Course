package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(substr("2754", 0, 4))
}

func Shannon(value string) float64 {
	var length = len(value)

	var frequency = make([]byte, 256)
	valueArr := strings.Split(value, "")
	for i := range valueArr {
		frequency[i]++
	}
	var entropy float64 = 0
	for i := 0; i < 256; i++ {
		var f = int(frequency[i])
		if f == 0 {
			continue
		}
		var k = float64(f / length)
		entropy += k * math.Log2(k) / math.Ln2
	}
	return entropy * (-1)
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
