package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("kitchen enterprises")
	http.HandleFunc("/api/sum/v1", sumHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	q := r.Form
	sum := sha256.Sum256([]byte(q.Get("shaQuery")))
	sha := fmt.Sprintf("%x", sum)
	w.Write([]byte(sha))
}
