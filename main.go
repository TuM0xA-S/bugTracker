package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type config struct {
	MongodbURI     string
	Host           string
	DBName         string
	CollectionName string
}



func main() {
	cfgFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(cfgFile)
	var cfg config
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongodbURI))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	bugs := client.Database(cfg.DBName).Collection(cfg.CollectionName)
	bugs.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{"CVE", 1},
			{"Source", 1},
		},
	})

	r := mux.NewRouter()
	r.Handle("/api/update", UpdateHandler{bugs})
	r.Handle("/api/cve/{CVE}", QueryHandler{bugs})
	http.Handle("/", r)
	if err := http.ListenAndServe(cfg.Host, nil); err != nil {
		log.Fatal(err)
	}
}
