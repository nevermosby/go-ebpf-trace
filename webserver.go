package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there from %s!", r.Host)
}

func main() {
	http.HandleFunc("/", handlerFunction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
