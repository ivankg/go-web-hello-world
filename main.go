package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/helloWorld", handler)
	serve := http.ListenAndServe(":80", nil)
	if serve != nil {
		log.Fatal("Listen and serve: ", serve)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world!")
}
