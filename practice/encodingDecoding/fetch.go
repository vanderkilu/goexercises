package main

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"
)

type Todo struct {
  userId string `json.userId`
  id int `json.userId`
  title string `json.title`
  completed string `json.completed`
}

type TodoList struct {
	results []Todo `json.results`
}

func main() {
	const url = "https://jsonplaceholder.typicode.com/todos/"

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		log.Fatalln("something went wrong fetching the url")
	}
	fmt.Println(response)

	var todoListResp TodoList
	err = json.NewDecoder(response.Body).Decode(&todoListResp)

	if err != nil {
		log.Fatalln("couldn't decode into todoListResp")
	}
	fmt.Println(todoListResp)

}