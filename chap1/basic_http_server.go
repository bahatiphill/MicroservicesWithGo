package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloworldResponse struct {
	Message string
}

func main() {

	PORT := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on PORT %v\n", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloworldResponse{Message: "Helloworld"}
	data, err := json.Marshal(response)

	if err != nil {
		panic("Oops, Can't encode the response to JSON")
	}
	fmt.Fprint(w, string(data))
}
