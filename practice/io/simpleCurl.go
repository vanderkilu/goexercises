package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
)


func main() {
	url := os.Args[1]
	fileName := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("couldn't fetch url")
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("couldn't read/create file")
		return
	}
	defer file.Close()
	dest := io.MultiWriter(os.Stdout, file)
	
	io.Copy(dest, resp.Body)

}