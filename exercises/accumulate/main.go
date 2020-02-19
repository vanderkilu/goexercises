package main

import (
	"strings"
	"fmt"
)

type Accumulator func(item string) string 

func accumulate(collection []string, fn Accumulator) []string {
	var results []string
	for _,item := range collection {
		results = append(results, fn(item))
	}
	return results
}

func main() {
	toLower := func(str string) string  {
		return strings.ToLower(str)
	}
	col := []string{"mAngo", "pineApple", "PAWPAW"}
	results := accumulate(col, toLower)
	fmt.Println(strings.Join(results, ","))
}