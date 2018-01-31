package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Message struct {
	Sender string `json:"txer"`
	Rec    string `json:"rxer"`
	Contet string `json:"msg"`
}

func maindis() {
	m := Message{
		Sender: "Alice",
		Rec:    "Bob",
		Contet: "Yo.",
	}

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(m); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
}
