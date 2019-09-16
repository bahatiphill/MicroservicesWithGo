package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type helloworldResponse struct {
	Message string `json:"message"`
}

type helloworldRequest struct {
	Name string `json:"name"`
}

func main() {

	PORT := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on PORT: %v\n", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	var request helloworldRequest
	err = json.Unmarshal(body, &request)

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response := helloworldResponse{Message: "Hello " + request.Name}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)

}
