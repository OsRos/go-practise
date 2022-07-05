package main

import (
	"html"
	"log"
	"net/http"
)

const BOOT_API_PREFIX = "http://localhost:8080"

func main() {
	http.HandleFunc("/command", handler)
	http.HandleFunc("/command-async", handler)
	http.HandleFunc("/command-async-annotation", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handler(writer http.ResponseWriter, request *http.Request) {
	requestPath := html.EscapeString(request.URL.Path)
	url := BOOT_API_PREFIX + requestPath
	go http.Get(url)
	/* 	if err != nil {
	   		log.Fatalf("Error while calling %s %v", url, err)
	   	}
	*/
}
