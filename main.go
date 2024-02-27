package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
	fmt.Println("server started")
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello, World!")
}
