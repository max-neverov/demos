package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	unmarshalHeader()
}

// START OMIT
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func unmarshalHeader() {
	b := []byte(`{"typ":"JWS","alg":"HS256","ALG":"none"}`)
	var h Header
	if err := json.Unmarshal(b, &h); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", h)
}

// END OMIT
