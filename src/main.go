package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiKey string `yaml:"key"`
}

type application struct {
	riotClient *golio.Client
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := golio.NewClient(
		config.ApiKey,
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New()),
	)

	var app application
	app.riotClient = client

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/champion", app.showChampion)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
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
