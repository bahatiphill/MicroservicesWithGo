package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	PORT := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on PORT %v\n", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!\n")
}
