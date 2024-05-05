package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /mongo", handlePostMongo)
	router.HandleFunc("POST /memstore", handlePostMemstore)
	router.HandleFunc("GET /memstore", handleGetMemstore)
}

func handlePostMongo(w http.ResponseWriter, r *http.Request) {
	// form fields are startDate, endDate, minCount, maxCount
}

func handlePostMemstore(w http.ResponseWriter, r *http.Request) {
	// form fields are key, value
}

func handleGetMemstore(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	fmt.Println("key: ", key)
}
