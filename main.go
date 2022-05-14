package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "bye")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello")
	})

	log.Println("server starting on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
