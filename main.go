package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type memObj struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
	return func(w http.ResponseWriter, r *http.Request) {
		var obj memObj
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			return
		}
		if err = json.Unmarshal(body, &obj); err != nil {
			fmt.Println("Invalid JSON data")
			return
		}
		memstore[obj.Key] = obj.Value
		w.Write([]byte(fmt.Sprintf("key: %s, value: %s", obj.Key, obj.Value)))

	}
}

func handleGetMemstore(memstore map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		value := memstore[key]
		w.Write([]byte(fmt.Sprintf("key: %s value: %s", key, value)))
	}
}
