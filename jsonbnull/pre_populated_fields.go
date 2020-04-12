package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	unmarshalUser()
}

// START OMIT
type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func unmarshalUser() {
	var v User

	json.Unmarshal([]byte(`{"name": "john"}`), &v)
	fmt.Println(v)

	json.Unmarshal([]byte(`{"surname": "smith"}`), &v)
	fmt.Println(v)
}

// END OMIT
