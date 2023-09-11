package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/freer4an/portfolio-website/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testStore *Store

func TestMain(m *testing.M) {
	cfg := util.InitConfig("../.")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.DBuri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	testStore = NewStore(client, cfg.DBname, "tests")
	_, err = testStore.collection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}

	exit := m.Run()
	defer os.Exit(exit)

	err = testStore.collection.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
