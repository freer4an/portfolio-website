package repository

import (
	"log"
	"os"
	"testing"

	"github.com/freer4an/portfolio-website/init/mongodb"
	"github.com/freer4an/portfolio-website/util"
	"go.mongodb.org/mongo-driver/mongo"
)

var testStore *Repository

func TestMain(m *testing.M) {
	cfg, err := util.InitConfig("../..")
	if err != nil {
		log.Fatal(err)
	}

	client, err := initMongo(cfg, m)
	if err != nil {
		log.Fatal(err)
	}

	testStore = New(client, cfg.DBname, "tests")

	exit := m.Run()
	defer os.Exit(exit)

	err = client.Database(cfg.DBname).Collection(cfg.CollName).Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func initMongo(cfg *util.Config, m *testing.M) (*mongo.Client, error) {
	client, err := mongodb.MongoClient(ctx, cfg)
	if err != nil {
		return nil, err
	}
	err = mongodb.MongoMigrate(ctx, cfg, client)
	if err != nil {
		return nil, err
	}
	return client, err
}
