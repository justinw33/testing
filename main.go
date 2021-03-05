package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("hello %s", time.Now().UTC().String()))
}
