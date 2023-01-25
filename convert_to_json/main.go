package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name  string
	Email string
}

func main() {
	users := []user{
		{
			Name:  "bassant",
			Email: "bassant.com",
		},
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		//return
	}

	fmt.Printf("json data: %s\n", jsonData)

	fmt.Printf("******************************")
	//fmt.Printf(users)
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func main() {
// 	data := map[string]interface{}{
// 		"intValue":    1234,
// 		"boolValue":   true,
// 		"stringValue": "hello!",
// 		"objectValue": map[string]interface{}{
// 			"arrayValue": []int{1, 2, 3, 4},
// 		},
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Printf("could not marshal json: %s\n", err)
// 		return
// 	}

// 	fmt.Printf("json data: %s\n", jsonData)
// }
