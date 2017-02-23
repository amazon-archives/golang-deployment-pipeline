package main

import (
	"io"
	"log"
	"net/http"
)

const version string = "2.0.1"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Current version: "+version)
}

func main() {
	log.Printf("Listening on port 8000...")
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8000", nil)
}
