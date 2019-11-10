package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	user := make(map[string]interface{})
	user["firstName"] = "enock"
	user["lastName"] = "kilu"
	user["address"] = map[string]interface{}{
		"location": "fijai-takoradi",
		"street": "fijai",
	}

	userToJson, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(userToJson))

}