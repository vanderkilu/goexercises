package main

import (
	"fmt"
	"math"
)

func toBase10(number int, base int) int {
	sum := 0
	pos := 0
	for number > 0 {
		rem := number % 10
		if number < 10 {
			sum += rem * int(math.Pow(float64(base), float64(pos)))
			break
		}
		number = number / 10
		sum += rem * int(math.Pow(float64(base), float64(pos)))
		pos += 1
	}
	return sum
}

func main() {
	fmt.Println(toBase10(1120, 3))
}
