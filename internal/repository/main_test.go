package repository

import (
	"log"
	"os"
	"testing"

	"github.com/freer4an/portfolio-website/inits/config"
	"github.com/freer4an/portfolio-website/inits/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

var testStore *Repository

const (
	test_collection = "tests"
)

func TestMain(m *testing.M) {
	cfg, err := config.InitConfig("../../configs")
	if err != nil {
		log.Fatal(err)
	}
	db_name := cfg.Database.Name

	client, err := initMongo(m, cfg.Database.Uri, db_name)
	if err != nil {
		log.Fatal(err)
	}

	testStore = New(client, db_name, test_collection)

	exit := m.Run()
	defer os.Exit(exit)

	err = client.Database(db_name).Collection("tests").Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func initMongo(m *testing.M, uri string, db_name string) (*mongo.Client, error) {
	client, err := mongodb.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}
	err = mongodb.MongoMigrate(client, db_name, test_collection)
	if err != nil {
		return nil, err
	}
	return client, err
}
