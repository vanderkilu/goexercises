package main
import (
	"strings"
	"bytes"
	"fmt"
)

/*
Acronym converts a string sentence into its acronym or 
abbreviated form. example Portable Network Graphics gives PNG
*/

func acronym(ac string) string {
	var acronym bytes.Buffer
	for _, word := range strings.Split(strings.TrimSpace(ac), " ")   {
		char := []rune(word)[0]
		acronym.WriteRune(char)
	}
	return strings.ToUpper(acronym.String())
}

func main() {
	ac := acronym("Three Letter Acronyms")
	fmt.Println(ac)
}