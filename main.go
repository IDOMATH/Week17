package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	memstore := make(map[string]string)

	router.HandleFunc("POST /mongo", handlePostMongo)
	router.HandleFunc("POST /memstore", handlePostMemstore)
	router.HandleFunc("GET /memstore", handleGetMemstore(memstore))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func handlePostMongo(w http.ResponseWriter, r *http.Request) {
	// form fields are startDate, endDate, minCount, maxCount
	w.Write([]byte("Posting to Mongo"))
}

func handlePostMemstore(w http.ResponseWriter, r *http.Request) {
	// form fields are key, value
	w.Write([]byte("Posting to memstore"))
}

func handleGetMemstore(memstore map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := memstore[key]
		w.Write([]byte(fmt.Sprintf("key: %s value: %s", key, value)))
	}
}
