package main

import (
	"bytes"
	"fmt"
	"log"
)

func maindis() {
	buf := &bytes.Buffer{}
	logger := log.New(buf, "demo time: ", log.LstdFlags)
	logger.Println("heelo, world")
	logger.Println("bye")
	fmt.Printf("%s\n", buf)
}
