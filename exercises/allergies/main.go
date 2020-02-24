package main

import (
	"fmt"
)

func allergy(score int) []string {
	var result []string
	allergies := map[string]int{
		"eggs":         1,
		"peanuts":      2,
		"shellfish ":   4,
		"strawberries": 8,
		"tomatoes":     16,
		"chocolate":    32,
		"pollen":       64,
		"cats":         128,
	}

	for k, v := range allergies {
		if score&v > 0 {
			result = append(result, k)
		}
	}
	return result
}

func main() {
	allegies := allergy(248)
	fmt.Println(allegies)
}
