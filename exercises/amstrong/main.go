package main

import (
	"fmt"
	"math"
)

func generateDigits(num int) []int {
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	return digits
}

func isAmstrong(num int) bool {
	digits := generateDigits(num)
	sum := 0
	for _, n := range digits {
		sum += int(math.Pow(float64(n), float64(len(digits))))
	}
	return sum == num
}

func main() {
	fmt.Println(isAmstrong(10))
}
