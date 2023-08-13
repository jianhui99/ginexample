package mongo

import (
	"context"
	"ginexample/config"
	"sync"

	"github.com/qiniu/qmgo"
)

var once sync.Once
var qmgoClient *qmgo.Client

func initMongo() {
	once.Do(func() {
		ctx := context.Background()
		newClient, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: getMongoURI()})
		if err != nil {
			panic(err)
		}
		qmgoClient = newClient
	})
}

func GetClient() *qmgo.Client {
	initMongo()
	return qmgoClient
}

func GetDatabase(dbName string) *qmgo.Database {
	return GetClient().Database(dbName)
}

func GetCollection(dbName string, collectionName string) *qmgo.Collection {
	return GetClient().Database(dbName).Collection(collectionName)
}

func getMongoURI() string {
	uri := config.GetEnv("MONGODB_URI")
	if uri == "" {
		return "mongodb://127.0.0.1:27017"
	}
	return uri
}
