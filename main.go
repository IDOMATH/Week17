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
	w.Write([]byte("Posting to Mongo"))
}

func handlePostMemstore(w http.ResponseWriter, r *http.Request) {
	// form fields are key, value
	w.Write([]byte("Posting to memstore"))
}

func handleGetMemstore(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	w.Write([]byte(fmt.Sprintf("key: %s", key)))
}
