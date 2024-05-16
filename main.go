package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
)

type memObj struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type mongoStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func newMongoStore(client *mongo.Client, dbName string) *mongoStore {
	return &mongoStore{
		client:     client,
		collection: client.Database(mongoDbName).Collection(mongoCollection),
	}
}

const portNumber = ":8080"
const dbUri = "mongodb://localhost:27017"
const mongoDbName = "week17"
const mongoCollection = "week17collection"

func main() {
	fmt.Println("Connecting to mongo")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	mongoStore := newMongoStore(client, mongoCollection)

	memstore := make(map[string]string)

	router.HandleFunc("POST /mongo", handlePostMongo(mongoStore))
	router.HandleFunc("POST /memstore", handlePostMemstore(memstore))
	router.HandleFunc("GET /memstore", handleGetMemstore(memstore))

	server := http.Server{
		Addr:    portNumber,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}

func handlePostMongo(store *mongoStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// form fields are startDate, endDate, minCount, maxCount
		w.Write([]byte("Posting to Mongo"))
	}
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
