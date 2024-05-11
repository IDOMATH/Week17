package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type memObj struct {
	key   string `json:"key"`
	value string `json:"value"`
}

func main() {
	router := http.NewServeMux()

	memstore := make(map[string]string)

	router.HandleFunc("POST /mongo", handlePostMongo)
	router.HandleFunc("POST /memstore", handlePostMemstore(memstore))
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

func handlePostMemstore(memstore map[string]string) http.HandlerFunc {
	// form fields are key, value
	return func(w http.ResponseWriter, r *http.Request) {
		var obj memObj
		err := json.NewDecoder(r.Body).Decode(&obj)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprintf("key: %s, value: %s", obj.key, obj.value)))
	}
}

func handleGetMemstore(memstore map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := memstore[key]
		w.Write([]byte(fmt.Sprintf("key: %s value: %s", key, value)))
	}
}
