package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/freer4an/portfolio-website/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db mongo.Database
var testStore *Store

func TestMain(m *testing.M) {
	cfg := util.InitConfig("../.")
	log.Println(cfg)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.DBuri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	testStore = NewStore(client, cfg.DBname, "tests")
	os.Exit(m.Run())

}
