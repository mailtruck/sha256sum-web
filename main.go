package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
