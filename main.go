package main
import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	fmt.Println("Hello world!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request ){ 

		fmt.Fprintf(w,"hello!")
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:8080",nil))


	
}

