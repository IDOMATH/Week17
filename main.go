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
	"time"
)

type memObj struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MongoRequest struct {
	StartDate string
	EndDate   string
	MinCount  int
	MaxCount  int
}

type MongoResponse struct {
	Code    int
	Msg     string
	Records []MongoRecord
}

type MongoRecord struct {
	Key        string
	CreatedAt  time.Time
	TotalCount int
}

type MongoStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoStore(client *mongo.Client, dbName string) *MongoStore {
	return &MongoStore{
		client:     client,
		collection: client.Database(mongoDbName).Collection(mongoCollection),
	}
}

func (s *MongoStore) InsertMongo(ctx context.Context, req MongoRequest) MongoResponse {

	res := MongoResponse{}
	return res
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

	mongoStore := NewMongoStore(client, mongoCollection)

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

func handlePostMongo(store *MongoStore) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// form fields are startDate, endDate, minCount, maxCount
		var mongoReq MongoRequest
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading request body")
			return
		}
		if err = json.Unmarshal(body, &mongoReq); err != nil {
			fmt.Println("Invalid JSON data")
			return
		}
		store.InsertMongo(context.Background(), mongoReq)
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
