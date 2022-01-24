package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/galazkamilosz/itemizer/src/model"
	"github.com/go-bongo/bongo"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiKey           string `yaml:"key"`
	ConnectionString string `yaml:"connectionString"`
	Database         string `yaml:"database"`
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	// mongo, err := connectToDatabase()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer mongo.Disconnect(context.Background())

	bon, err := connectDb()
	if err != nil {
		log.Fatal(err)
	}

	client := golio.NewClient(
		config.ApiKey,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New()),
	)
	champions, err := client.DataDragon.GetChampions()
	if err != nil {
		log.Fatal(err)
	}
	// collection := mongo.Database("itemizer").Collection("champions")
	for _, champion := range champions {
		champ := model.Champion{
			Champ: champion,
		}
		bon.Collection("champions").Save(&champ)
	}
}

func getConfig() (*Config, error) {
	config := &Config{}
	file, err := os.Open("./private/config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func connectToDatabase() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func connectDb() (*bongo.Connection, error) {
	connection, err := bongo.Connect(&bongo.Config{
		ConnectionString: "localhost",
		Database:         "itemizer",
	})
	if err != nil {
		return nil, err
	}
	return connection, nil
}
