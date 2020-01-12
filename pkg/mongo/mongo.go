package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

//mongodb配置
var (
	Ctx           context.Context
	client        *mongo.Client
	LogCollection *mongo.Collection
)

func init() {
	var err error
	Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	client, err = mongo.Connect(Ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("cant't connect mongo server.", err)
	}
	logcon()
}

func logcon() {
	LogCollection = client.Database("w3fy").Collection("logs")
}
