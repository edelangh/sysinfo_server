
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !")
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0.0.0")
}

func handlerDuration(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "0s")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/version", handlerVersion)
	http.HandleFunc("/duration", handlerDuration)

	log.Print("Start by visiting http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

