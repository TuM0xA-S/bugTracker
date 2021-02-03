package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/TuM0xA-S/bugTracker/load/redhat"
	"github.com/TuM0xA-S/bugTracker/load/ubuntu"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/TuM0xA-S/bugTracker/load/debian"
	"github.com/TuM0xA-S/bugTracker/types"
	"go.mongodb.org/mongo-driver/mongo"
)

//UpdateHandler updates his collection
type UpdateHandler struct {
	bugs *mongo.Collection
}

func (u UpdateHandler) pushToDB(list []types.BugData) {
	for _, v := range list {
		_, err := u.bugs.ReplaceOne(context.TODO(), bson.M{
			"Source": v.Source,
			"CVE":    v.CVE,
		},
			v,
			options.Replace().SetUpsert(true),
		)
		if err != nil {
			log.Printf("debian troubles when update: %d\n", err)
		}
	}
}

func (u UpdateHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("when update: %v", err)
		}
	}()
	log.Println("update started...")
	start := time.Now()
	list, err := debian.Load()
	if err != nil {
		log.Printf("debian update failed: %v\n", err)
	}
	u.pushToDB(list)

	list, err = ubuntu.Load()
	if err != nil {
		log.Printf("ubuntu update failed: %v\n", err)
	}
	u.pushToDB(list)

	list, err = redhat.Load()
	if err != nil {
		log.Printf("redhat update failed: %v\n", err)
	}
	u.pushToDB(list)

	rw.Write([]byte("\"update done\"\n"))
	log.Printf("update ended. time elapsed %fs", time.Since(start).Seconds())
}

//QueryHandler handles queries
type QueryHandler struct {
	bugs *mongo.Collection
}

func (q QueryHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println("query...")
	CVE := "CVE-" + mux.Vars(req)["CVE"]
	query := bson.D {
		{"CVE", CVE},
	}
	qs := req.URL.Query()
	if qs.Get("source") != "" {
		query = append(query, bson.E{"Source", qs.Get("source")})
	}
	if qs.Get("pkg") != "" {
		query = append(query, bson.E{"Packages", qs.Get("pkg")})
	}
	cur, err := q.bugs.Find(context.TODO(), query)
	if err != nil {
		log.Printf("on query: %v\n", err)
	}
	var bd []types.BugData
	if err := cur.All(context.TODO(), &bd); err != nil {
		log.Printf("when decoding result from bd: %v\n", err)
	}
	
	enc := json.NewEncoder(rw)
	enc.SetIndent("", "    ")
	if err := enc.Encode(bd); err != nil {
		log.Printf("when encoding result to json: %v\n", err)
	}
}
