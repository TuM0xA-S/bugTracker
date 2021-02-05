package main

import (
	"context"
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
	Indent         string
}

func parseConfig() (cfg config) {
	//set default
	cfg.MongodbURI = "mongodb://localhost:27017"
	cfg.Host = "localhost:2222"
	cfg.DBName = "bugTracker"
	cfg.CollectionName = "bugs"
	cfg.Indent = ""

	//read cfg from env
	if val, ok := os.LookupEnv("BT_MongodbURI"); ok {
		cfg.MongodbURI = val
	}
	if val, ok := os.LookupEnv("BT_Host"); ok {
		cfg.Host = val
	}
	if val, ok := os.LookupEnv("BT_DBName"); ok {
		cfg.DBName = val
	}
	if val, ok := os.LookupEnv("BT_CollectionName"); ok {
		cfg.CollectionName = val
	}
	if val, ok := os.LookupEnv("BT_Indent"); ok {
		cfg.Indent = val
	}
	return
}

func main() {
	log.Println("starting...")

	cfg := parseConfig()
	log.Printf("%#v\n", cfg)
	ctx := context.Background()

	log.Println("connecting to db on", cfg.MongodbURI)
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
	if _, err := bugs.Indexes().CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{Key: "CVE", Value: 1},
			{Key: "Source", Value: 1},
		},
	}); err != nil {
		log.Println("while trying to create index:", err)
	}

	r := mux.NewRouter()
	r.Handle("/api/update", UpdateHandler{bugs})
	r.Handle("/api/cve/{CVE}", QueryHandler{bugs, cfg.Indent})
	http.Handle("/", r)

	log.Println("starting serving at", cfg.Host)
	if err := http.ListenAndServe(cfg.Host, nil); err != nil {
		log.Fatal(err)
	}
}
