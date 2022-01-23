package main

import (
	"fmt"
	"log"
	"os"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiKey string `yaml:"key"`
}

func main() {
	config := Config{}
	file, err := os.Open("./private/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	client := golio.NewClient(
		config.ApiKey,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New()),
	)
	summoner, err := client.Riot.LoL.League.Get("GattoRosso")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(summoner.Name)
}
