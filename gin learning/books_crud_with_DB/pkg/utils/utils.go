//every thing we will recieve from user will be as json and evey thing we need to give it to the user
//will be as json file
//so when we deal with our code we need to convert json to be able to deal with in our controller function
// this called marshell and ummarshell the data

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// create fuction to unmarshell body when you get from users
func ParseBody(r *http.Request, X interface{}) { // r is used to acces the request that came from user
	//read the body and if there is no error start to umnarshell the body
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), X); err != nil {
			//if there is error
			return
		}
	}
} // func
