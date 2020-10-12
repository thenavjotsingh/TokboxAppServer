package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/get" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "server started!!!")

	token := createToken()
	respSession := createSession(token)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respSession)

	// fmt.Fprintf(w, respSession)

}

func main() {
	http.HandleFunc("/get", getHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
