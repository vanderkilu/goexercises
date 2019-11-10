package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type (
	Todo struct {
		UserId int `json:"userId"`
		Id int `json:"id"`
		Title string `json:"title"`
		Completed bool `json:"completed"`
	}
)

func main() {
	const url = "https://jsonplaceholder.typicode.com/todos/"

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("something went wrong fetching the url")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("something went wrong fetching the url")
	}

	defer response.Body.Close()

	var todoListResp []Todo
	err = json.Unmarshal(body, &todoListResp)

	if err != nil {
		log.Fatalln(err)
	}
	
	for _, item := range todoListResp {
		fmt.Printf("{\n")
		fmt.Printf("\tid: %d\n", item.Id)
		fmt.Printf("\ttitle: %s\n", item.Title)
		fmt.Printf("\tcompleted: %v\n", item.Completed)
		fmt.Printf("}\n")
	}

}