package main

import (
	"encoding/json"
)

type user struct {
	Name  string
	Email string
}

func main() {
	users := []user{
		{
			Name:  "alaa",
			Email: "alaa.attya91@gmai.com",
		},
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		println("could not marshal json: %s\n", err)
		return
	}

	println("json data: \n", jsonData)

	println("******************************")
	println(users)
}
