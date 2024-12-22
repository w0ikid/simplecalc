package main

import (
	"log"
	"net/http"
	"github.com/w0ikid/simplecalc/internal/api"
)

func main() {
	http.HandleFunc("/api/v1/calculate", api.CalculateHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
